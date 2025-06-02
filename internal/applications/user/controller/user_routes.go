package controller

import (
	"hokusai/internal/component/log"

	"github.com/labstack/echo/v4"
)

func (c *UserController) AddRoutes(e *echo.Echo, appName string) {
	group := e.Group(appName + "/user")
	group.Use(c.authService.ValidateClientKeyMiddleware())
	applyLoggerMiddleware(group)

	group.POST("", c.Create)
	group.PUT("/:id", c.Update)
	group.DELETE("/:id", c.Delete)
	group.GET("/:id", c.GetById)
	group.GET("", c.GetAll)
	group.POST("/login", c.Login)

}

// LOGGER
func applyLoggerMiddleware(group *echo.Group) {
	group.Use(log.ContextualLoggerMiddleware(log.ContextualLoggerConfig{
		DomainName: "user-account",
	}))
}
