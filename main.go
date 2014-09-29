package main

import (
	"github.com/codegangsta/cli"
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
	port := c.String(FieldPort)
	err := http.ListenAndServe(":"+port, http.FileServer(http.Dir(".")))
	if err != nil {
		log.Fatalln(err)
	}
}
