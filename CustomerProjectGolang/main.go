package main

import (
	"github.com/joho/godotenv"
	"github.com/rushikeshkandekar/controller"
	"github.com/rushikeshkandekar/router"
	"log"
)

func main() {

	log.Println("Started the server...........")

	err := godotenv.Load("app.env")
	controller.Handleerror(err)

	router.Router()

	//err2 := database.ConnectDB()
	//defer err2.Close()

}
