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
	// helpers helpers.App
)

func main() {

	fmt.Println("Starting the application...")
	port := "3010"
	serverIP := "localhost"

	helpers.DBAccess(false, "dev")

	server := negroni.Classic()
	mux := httprouter.New()

	// NOTE: prevent sending of Call Stack on panic
	recovery := negroni.NewRecovery()
	recovery.PrintStack = false
	server.Use(recovery)

	server.Use(negroni.HandlerFunc(handlers.CORS))

	server.UseHandler(mux)

	// a := helpers.App{}
	// a.Initialize("root", "root", "ums_db")

	mux.GET("/users", users.GetUsers)
	mux.GET("/user/:id", users.GetUser)
	mux.POST("/user", users.CreateUser)
	mux.DELETE("/user/:id", users.DeleteUser)
	mux.PUT("/user/:id", users.UpdateUser)

	server.Run("localhost:" + port)
	graceful.Run(serverIP+":"+port, 10*time.Second, server)
}
