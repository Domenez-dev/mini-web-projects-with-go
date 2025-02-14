package main

import (
	"fmt"
	"net/http"
	"time"
)

type Login struct {
	hashedPassword string
	sessionToken   string
	CSFTToken      string
}

// users data: {username: Login}
var users = map[string]Login{}

func main() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/register", register)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/private", private)
	http.ListenAndServe(":8080", nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		err := http.StatusMethodNotAllowed
		http.Error(w, "Invalide Methode", err)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	user, ok := users[username]
	if !ok || !CheckPasswordHash(password, user.hashedPassword) {
		err := http.StatusUnauthorized
		http.Error(w, "User not found", err)
		return
	}

	sessionToken := GenerateToke(32)
	csrfToken := GenerateToke(32)
	oneDay := time.Now().Add(24 * time.Hour)

	// set Session cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Expires:  oneDay,
		HttpOnly: true,
	})

	// set CSRF cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    csrfToken,
		Expires:  oneDay,
		HttpOnly: false,
	})

	user.sessionToken = sessionToken
	user.CSFTToken = csrfToken
	users[username] = user

	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintf(w, "User Logged succesfully")
}

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		err := http.StatusMethodNotAllowed
		http.Error(w, "Invalide Methode", err)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Check if the username and password are valid
	if !StrongPassword(password) || len(username) < 8 {
		err := http.StatusNotAcceptable
		http.Error(w, "Invalide username or password", err)
		return
	}

	// Check If user exists
	if _, ok := users[username]; ok {
		err := http.StatusConflict
		http.Error(w, "User already exists", err)
		return
	}

	HashPassword, _ := HashPassword(password)
	users[username] = Login{hashedPassword: HashPassword}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User created")
}

func logout(w http.ResponseWriter, r *http.Request) {
	if Authorize(r) != nil {
		err := http.StatusUnauthorized
		http.Error(w, "UnAuthorized", err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: false,
	})

	username := r.FormValue("username")
	user := users[username]
	user.sessionToken = ""
	user.CSFTToken = ""
	users[username] = user

	fmt.Fprintf(w, "Logged out succesfully")
}

func private(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		err := http.StatusUnauthorized
		http.Error(w, "Unauthorized method", err)
		return
	}

	if Authorize(r) != nil {
		err := http.StatusUnauthorized
		http.Error(w, "Invalide Token", err)
		return
	}
	username := r.FormValue("username")
	fmt.Fprintf(w, "CSRF valide, welcome back %s", username)
}
