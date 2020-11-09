package server

import (
	"time"

	"github.com/SKilliu/taxi-service/config"
	"github.com/SKilliu/taxi-service/server/middlewares"

	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const durationThreshold = time.Second * 10

// Router for app
func Router(cfg config.Config) *echo.Echo {
	router := echo.New()

	cors := middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*", "GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"*", "Accept", "Authorization", "Content-Type", "X-CSRF-Token", "x-auth", "Access-Control-Allow-Origin", "Access-Control-Allow-Methods", "Access-Control-Allow-Credentials"},
		ExposeHeaders:    []string{"*", "Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	provider := NewProvider(cfg, cfg.DB())
	m := middlewares.New(cfg)

	router.Use(
		cors,
		middleware.Recover(),
		middleware.LoggerWithConfig(middleware.DefaultLoggerConfig),
	)

	router.GET("/swagger/*", echoSwagger.WrapHandler)

	router.POST("/sign_up", provider.AuthHandler.SignUp)

	authGroup := router.Group("")
	authGroup.Use(m.ParseToken)

	router.POST("/sign_in", provider.AuthHandler.SignIn)

	router.POST("/user", provider.UsersHandler.EditProfile)
	router.GET("/user", provider.UsersHandler.GetProfile)

	return router
}
