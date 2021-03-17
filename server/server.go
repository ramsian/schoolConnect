package server

import (
	"main/routes"
)

func Init() {
	router := routes.Init()
	router.Run()
}
