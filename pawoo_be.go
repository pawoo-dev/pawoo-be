package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	r *gin.Engine
)

func main() {
	r = gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*", "http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE", "OPTIONS", "HEAD"},
		AllowHeaders:     []string{"authorization", "access-control-allow-origin", "Access-Control-Allow-Headers", "Origin", "Content-Length", "Content-Type", "authentication", "Access-Control-Request-Method", "Access-Control-Request-Headers"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome to pawoo!")
		return
	})
	if err := r.Run(":8080"); err != nil {
		logrus.WithField("error", err).Errorf("http server failed to start")
	}
}
