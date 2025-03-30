package main

import (
	"net/http"
	"os"

	"github.com/chaso-pa/ecomimic-cron/middleware"
	"github.com/chaso-pa/ecomimic-cron/services"
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

	c.AddFunc("32 */2 * * *", func() { services.CrawlAllArticles() })

	c.Start()

	r.GET("/get_article", func(ctx *gin.Context) {
		err := services.CrawlAllArticles()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	if os.Getenv("GIN_MODE") != "release" {
		r.Run("localhost:" + port)
	} else {
		r.Run(":" + port)
	}
}
