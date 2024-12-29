package main

import "example/jwt-authentication-tutorial-project/database"

func main() {
	database.Connect("root:root@tcp(localhost:3306)/jwt_demo?parseTime=true")
}