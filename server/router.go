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

	// Users handlers
	router.GET("/user", provider.UsersHandler.GetProfile)
	router.POST("/user", provider.UsersHandler.EditProfile)
	router.DELETE("/user", provider.UsersHandler.Delete)
	router.PATCH("/user", provider.UsersHandler.UpdateProfileImage)

	// Operators handlers
	router.POST("/operators/user", provider.OperatorsHandler.CreateNewUser)
	router.POST("/operators/car", provider.OperatorsHandler.AddNewCar)
	router.POST("/operators/assign", provider.OperatorsHandler.AssignCarToDriver)

	// Orders handler
	router.POST("/orders", provider.OrdersHandler.CreateOrder)
	router.GET("/orders", provider.OrdersHandler.GetAvailableOrders)
	router.PATCH("/orders", provider.OrdersHandler.OrderActions)

	return router
}
