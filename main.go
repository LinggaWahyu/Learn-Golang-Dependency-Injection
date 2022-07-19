package main

import (
	"learn-golang-restful-api/helper"
	"learn-golang-restful-api/middleware"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}
}

func main() {

	server := InitializedServer()

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
