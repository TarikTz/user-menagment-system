package main

import (
	"fmt"
	"time"

	"github.com/TarikTz/user-menagment-system/controllers"
	"github.com/TarikTz/user-menagment-system/helpers"
	"github.com/julienschmidt/httprouter"
	"github.com/tylerb/graceful"
	"github.com/urfave/negroni"
)

var (
	users controllers.Users
	handlers helpers.Handlers
)

func main() {

	fmt.Println("Starting the application...")
	port := "3010"
	serverIP := "localhost"

	helpers.DBAccess()

	server := negroni.Classic()
	mux := httprouter.New()

	recovery := negroni.NewRecovery()
	recovery.PrintStack = false
	server.Use(recovery)

	server.Use(negroni.HandlerFunc(handlers.CORS))

	server.UseHandler(mux)

	mux.GET("/users", users.GetUsers)
	mux.GET("/user/:id", users.GetUser)
	mux.POST("/user", users.CreateUser)
	mux.DELETE("/user/:id", users.DeleteUser)
	mux.PUT("/user/:id", users.UpdateUser)

	mux.GET("/", users.RenderAccess)

	server.Run("0.0.0.0:" + port)
	graceful.Run(serverIP+":"+port, 10*time.Second, server)
}