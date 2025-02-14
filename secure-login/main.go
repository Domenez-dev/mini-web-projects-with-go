package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Login struct {
	Username       string
	HashedPassword string
	SessionToken   string
	CSRFToken      string
}

var DB *sql.DB

func main() {
	InitDatabase()
	defer DB.Close()

	http.HandleFunc("/login", login)
	http.HandleFunc("/register", register)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/private", private)
	http.ListenAndServe(":8080", nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	user, err := GetUser(username)
	if err != nil || !CheckPasswordHash(password, user.HashedPassword) {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	sessionToken := GenerateToken(32)
	csrfToken := GenerateToken(32)
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

	user.SessionToken = sessionToken
	user.CSRFToken = csrfToken
	if err := UpdateUser(user); err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintf(w, "User Logged in successfully")
}

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Check if the username and password are valid
	if !StrongPassword(password) || len(username) < 8 {
		http.Error(w, "Invalid username or password", http.StatusNotAcceptable)
		return
	}

	// Check if user exists
	if _, err := GetUser(username); err == nil {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}

	hashedPassword, _ := HashPassword(password)
	user := Login{Username: username, HashedPassword: hashedPassword}
	if err := CreateUser(user); err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User created")
}

func logout(w http.ResponseWriter, r *http.Request) {
	if err := Authorize(r); err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Clear cookies
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
	user, err := GetUser(username)
	if err != nil {
		http.Error(w, "yFailed to get user", http.StatusInternalServerError)
		return
	}

	user.SessionToken = ""
	user.CSRFToken = ""
	if err := UpdateUser(user); err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Logged out successfully")
}

func private(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Unauthorized method", http.StatusUnauthorized)
		return
	}

	if err := Authorize(r); err != nil {
		http.Error(w, "Invalid Token", http.StatusUnauthorized)
		return
	}
	username := r.FormValue("username")
	fmt.Fprintf(w, "CSRF valid, welcome back %s", username)
}
