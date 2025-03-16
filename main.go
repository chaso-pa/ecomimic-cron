package main

import (
	"os"

	"github.com/chaso-pa/ecomimic-cron/middleware"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

const defaultPort = "8080"

func main() {
	middleware.LoadEnv()
	middleware.ConDb()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	r := gin.Default()
	c := cron.New()

	c.Start()

	if os.Getenv("GIN_MODE") != "release" {
		r.Run("localhost:" + port)
	} else {
		r.Run(":" + port)
	}
}
