package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/gops/agent"
	"log"
	"net/http"
)

var db = make(map[string]string)

func globalRecover(c *gin.Context) {
	defer func(c *gin.Context) {
		if rec := recover(); rec != nil {
			fmt.Println()
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": rec,
			})
		}
	}(c)
	c.Next()
}

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(globalRecover)

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.AbortWithError(404,errors.New("not found"))
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//        "foo":  "bar",
	//        "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080


	if err := agent.Listen(agent.Options{}); err != nil {
		log.Fatalf("agent.Listen err: %v", err)
	}

	r.Run(":8080")
}