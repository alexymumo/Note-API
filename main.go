package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alexymumo/controllers"
	"github.com/joho/godotenv"
)

var server = controllers.Server{}

func main() {
	//var err error
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error occurred getting env, %v", err)
	} else {
		fmt.Printf("Success getting env")
	}
	//Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string

	server.InitializeDb(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	server.Run(":8000")

}

/*
DB_HOST=127.0.0.1
DB_DRIVER=mysql
API_SECRET=2323232FAGADSAGDAGAS
DB_PASSWORD=
DB_USER=
DB_NAME=notedb
DB_PORT=3306*/
