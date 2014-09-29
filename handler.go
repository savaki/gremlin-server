package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func RequireSsl(c *gin.Context) {
	if proto := c.Request.Header.Get("X-Forwarded-Proto"); proto != "" && proto != "https" {
		c.Redirect(302, fmt.Sprintf("https://%s%s", c.Request.Host, c.Request.RequestURI))
		c.Abort(302)
		return
	}

	c.Next()
}
