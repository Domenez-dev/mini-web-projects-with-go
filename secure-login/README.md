# GoLang Web Project

## Description

This is a simple web project written in GoLang that demonstrates user authentication mechanisms including login, registration, CSRF protection, and session management using cookies. The project uses SQLite as the database.

## Project Structure

- `main.go`: Entry point of the application. Sets up HTTP routes and starts the server.
- `database.go`: Contains functions for initializing the SQLite database.
- `utils.go`: Utility functions for password hashing, token generation, and password strength validation.
- `middleware.go`: Middleware functions for authorization.
- `services.go`: Service functions for interacting with the database.

## Endpoints

- **POST /register**: Registers a new user.
- **POST /login**: Logs in an existing user.
- **POST /logout**: Logs out the user by clearing the session and CSRF tokens.
- **POST /private**: Accesses a private endpoint that requires CSRF token validation.

## Running the Project

1. Install GoLang and SQLite.
2. Clone the repository.
3. Navigate to the project directory.
4. Run `go run main.go database.go utils.go middleware.go services.go` to start the server.

## Testing the Project

A bash script `test_scripts.sh` is provided to test the endpoints using `curl`.

To execute the tests:

1. Ensure the server is running.
2. Run the bash script: `./test_scripts.sh`

## Dependencies

- `github.com/mattn/go-sqlite3`
- `golang.org/x/crypto/bcrypt`

## Security Measures

- Passwords are hashed using bcrypt.
- CSRF protection is implemented using tokens.
- Session management is handled using secure cookies.

## Notes

- Ensure the database file `awesome.db` is in the project directory.
- Modify the `BASE_URL` in `test_scripts.sh` if the server runs on a different host or port.
