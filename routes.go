package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/savaki/stormpath-go/accounts"
	"github.com/savaki/stormpath-go/auth"
	"log"
	"os"
)

func fail(c *gin.Context, err error) {
	c.JSON(500, map[string]string{"status": err.Error()})
}

func AddStaticRoutes(routes *gin.Engine) {
	routes.Static("/js", "public/js")
	routes.Static("/public", "public")
}

func AddOauthRoutes(routes *gin.Engine) {
	apiKey, err := auth.EnvAuth()
	if err != nil {
		log.Fatalln(err)
	}
	api := accounts.FromUrl(apiKey, os.Getenv("STORMPATH_URL"))

	routes.GET("/oauth2callback", func(c *gin.Context) {
		query := struct {
			AccessCode string `json:"access_code" form:"access_code"`
		}{}
		if c.Bind(&query) {
			account, err := api.LoginViaGoogle(query.AccessCode)
			if err != nil {
				fail(c, err)
			} else {
				c.JSON(200, account)
			}
		} else {
			fail(c, fmt.Errorf("unable to do stuff"))
		}
	})
}
