package main

import "example/jwt-authentication-tutorial-project/database"

func main() {
	database.Connect("postgresql://postgres:postgres@localhost:5432/jwtautherntication")
}