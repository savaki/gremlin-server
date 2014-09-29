package main

import (
	"github.com/gin-gonic/gin"
)

func AddStaticRoutes(routes *gin.Engine) {
	routes.Static("/", "public")
}
