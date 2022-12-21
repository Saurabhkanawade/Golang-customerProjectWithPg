package main

import (
	"github.com/Saurabhkanawade/golang-practice/helper"
	"github.com/Saurabhkanawade/golang-practice/router"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting the server on 9000.....................")

	err := godotenv.Load("app.env")
	helper.CheckErrorNill(err)

	log.Println("godotenv log in main()............")

	route := router.Router()

	// Starting Server
	log.Fatal(http.ListenAndServe(":8000", route))

	//db := database.Connect()
	//defer db.Close()
}
