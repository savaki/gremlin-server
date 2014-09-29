package main

import (
	"github.com/codegangsta/cli"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

const (
	FieldPort = "port"
)

var (
	FlagPort = cli.StringFlag{FieldPort, "8080", "the port to use", "PORT"}
)

func main() {
	app := cli.NewApp()
	app.Name = "gremlin-server"
	app.Usage = "webserver for gremlin service"
	app.Author = "Matt Ho matt.ho@gmail.com"
	app.Flags = []cli.Flag{
		FlagPort,
	}
	app.Action = server

	app.Run(os.Args)
}

func server(c *cli.Context) {
	routes := gin.Default()
	routes.Use(RequireSsl)
	AddStaticRoutes(routes)

	port := c.String(FieldPort)
	err := http.ListenAndServe(":"+port, routes)
	if err != nil {
		log.Fatalln(err)
	}
}
