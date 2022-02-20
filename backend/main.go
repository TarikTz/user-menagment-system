package main

import (
	"fmt"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/tylerb/graceful"
	"github.com/urfave/negroni"
)

func main() {

	fmt.Println("Starting the application...")
	port := "3090"
	serverIP := "localhost"

	server := negroni.Classic()
	mux := httprouter.New()

	// NOTE: prevent sending of Call Stack on panic
	recovery := negroni.NewRecovery()
	recovery.PrintStack = false
	server.Use(recovery)

	server.UseHandler(mux)

	server.Run("localhost:" + port)
	graceful.Run(serverIP+":"+port, 10*time.Second, server)
}
