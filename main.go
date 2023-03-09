package main

import (
	_ "github.com/go-sql-driver/mysql"
	"golang-crud/controller"
)

var server = controller.Server{}

func main() {

	server.Initialize()

	server.Run(":8080")
}
