package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alexymumo/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) InitializeDb(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error
	if Dbdriver == "mysql" {
		dburl := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbName)
		server.DB, err = gorm.Open(Dbdriver, dburl)
		if err != nil {
			log.Fatal("An error occurred", err)
		} else {
			fmt.Printf("Connected to %s successfully", Dbdriver)
		}
		server.DB.Debug().AutoMigrate(&models.Note{})

		server.Router = mux.NewRouter()

		server.initializeRoutes()
	}
}

func (server *Server) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, server.Router))

}
