package main

import (
	"fmt"
	"main/database"
	"main/server"
)

func main() {
	fmt.Println("It works!")
	database.ConnectDatabase()
	server.Init()
}
