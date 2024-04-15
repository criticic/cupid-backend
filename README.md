# cupid-backend

backend for [cupid](https://github.com/criticic/cupid-frontend/tree/main), written in go using go-fiber & firebase admin sdk.

Note: Requires is Service Account JSON file from Firebase.

## File Structure

- `main.go` - entry point, and contains api route locations

- `utils/`
  - `firebase.go` - contains firebase admin sdk setup

- `random_name_generator/` pkg contains the functions that generate the username

- `controllwers/` contains the controllers for the api routes

## Running

- `go get` to install dependencies
- `go run main.go` to start the server
