package main

func GetUser(username string) (Login, error) {
	var user Login
	row := DB.QueryRow("SELECT username, hashedPassword, sessionToken, csrfToken FROM users WHERE username = ?", username)
	err := row.Scan(&user.Username, &user.HashedPassword, &user.SessionToken, &user.CSRFToken)
	return user, err
}

func CreateUser(user Login) error {
	_, err := DB.Exec("INSERT INTO users (username, hashedPassword, sessionToken, csrfToken) VALUES (?, ?, ?, ?)",
		user.Username, user.HashedPassword, user.SessionToken, user.CSRFToken)
	return err
}

func UpdateUser(user Login) error {
	_, err := DB.Exec("UPDATE users SET hashedPassword = ?, sessionToken = ?, csrfToken = ? WHERE username = ?",
		user.HashedPassword, user.SessionToken, user.CSRFToken, user.Username)
	return err
}
