package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pawoo-dev/pawoo-be/controller"
	"github.com/pawoo-dev/pawoo-be/dao"
	"github.com/pawoo-dev/pawoo-be/handler"
	"github.com/pawoo-dev/pawoo-be/helper"
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
	var hostname string
	user, pass := helper.GetDatabaseSecrets()
	hostname = user + ":" + pass + "@tcp(pawoo-db.cveaywscwq9w.ap-southeast-1.rds.amazonaws.com:3306)/pawoo?parseTime=true&charset=utf8mb4"
	//hostname = configObj.Dsn

	// init db
	err := dao.InitDB(hostname)
	if err != nil {
		panic(err)
	}

	// init controller
	controller.NewAuthenticationController()
	controller.NewUserController()

	// init handler
	registerHandler()

	// run http server
	if err := r.Run(":8080"); err != nil {
		logrus.WithField("error", err).Errorf("http server failed to start")
	}
}

func registerHandler() {
	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome to pawoo!")
		return
	})
	v1 := r.Group("/api/v1")
	{
		v1.POST("/login", handler.LoginHandler)
		v1.POST("/signup", handler.SignUpHandler)
		v1.POST("/confirm", handler.ConfirmUserHandler)
		v1.POST("/resend", handler.ResendChallengeCodeHandler)
	}
	protectedV1 := r.Group("/api/v1")
	protectedV1.Use(handler.AuthMiddlewareHandler)
	{
		// authentication handler
		protectedV1.POST("/logout", handler.LogoutUserHandler)
	}
}
