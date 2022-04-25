package main

import (
	"database/sql"
	"fmt"
	"resto/handler"
	"resto/repository"
	"resto/server"
	"resto/service"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	user      = "postgres"
	password  = "handoko"
	dbname    = "resto"
	port      = 5432
	ssl       = "disable"
	time      = "Asia/Jakarta"
	localhost = "localhost:5000"
)

func main() {
	psqlInfo := fmt.Sprintf("user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		user, password, dbname, port, ssl, time)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	repository := repository.Repository{
		DB: db,
	}

	service := service.Service{
		Repository: repository,
	}

	handler := handler.Handler{
		Service: service,
	}

	router := gin.Default()

	start := server.Server{
		Router:  router,
		Handler: &handler,
	}
	start.StartServer()
	router.Run(localhost)

}
