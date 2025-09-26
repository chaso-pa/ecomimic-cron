package main

import (
	"log"
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
	c.AddFunc("*/5 * * * *", func() { services.CrawlEconomicCalendar() })

	c.Start()

	r.GET("/get_article", func(ctx *gin.Context) {
		err := services.CrawlAllArticles()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	r.GET("/test", func(c *gin.Context) {
		err := services.CrawlEconomicCalendar()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
