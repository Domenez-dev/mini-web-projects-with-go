package main

import (
	"crypto/rand"
	"encoding/base64"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func StrongPassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	hasUpper := false
	hasLower := false
	hasNumber := false

	for _, char := range password {
		if char >= 'A' && char <= 'Z' {
			hasUpper = true
		}

		if char >= 'a' && char <= 'z' {
			hasLower = true
		}

		if char >= '0' && char <= '9' {
			hasNumber = true
		}
	}

	return hasUpper && hasLower && hasNumber
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateToke(length int) string {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		log.Fatal("Failed to generate token")
	}
	return base64.URLEncoding.EncodeToString(bytes)
}
