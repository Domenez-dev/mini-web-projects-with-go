package main

import (
	"errors"
	"fmt"
	"net/http"
)

var AuthError = errors.New("UnAuthorized")

func Authorize(r *http.Request) error {
	username := r.FormValue("username")
	user, err := GetUser(username)
	if err != nil {
		return AuthError
	}

	session_token, err := r.Cookie("session_token")

	fmt.Println("user token: ", user.CSRFToken)
	fmt.Println("provided_t: ", session_token)
	if err != nil || session_token.Value == "" || session_token.Value != user.SessionToken {
		return AuthError
	}

	csrf_token := r.Header.Get("X-CSRF-Token")
	if csrf_token != user.CSRFToken || csrf_token == "" {
		return AuthError
	}

	return nil
}
