package main

//TODO:
/*
	initialize zap for logs in main and initialize
	Add tests
	login
	docker

Consume data from external API

	Use IMROC package for API consumption
*/

import (

	// "go.uber.org/zap"

	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jerry/message-hub-template-go/template/dig"
	"github.com/jerry/message-hub-template-go/template/server"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	}
}

func run() error {
	g := gin.Default()
	d := dig.BuildProject()
	svr := server.NewServer(g, d)
	svr.MapRoutes()
	if err := svr.SetupDB(); err != nil {
		return err
	}
	return svr.Start()
}
