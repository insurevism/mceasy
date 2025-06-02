package controller

import (
	"mceasy/internal/applications/auth/dto"
	"mceasy/internal/applications/auth/service"
	"mceasy/internal/helper"
	"mceasy/internal/helper/response"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	service service.AuthService
}

func NewAuthController(service service.AuthService) *AuthController {
	return &AuthController{service}
}

func (c *AuthController) AddRoutes(e *echo.Echo, appName string) {
	group := e.Group(appName + "/auth")

	group.POST("/register", c.Register)
	group.POST("/login", c.Login)
}

func (c *AuthController) Register(ctx echo.Context) error {
	payload := &dto.RegisterRequest{}
	err := helper.BindAndValidate(ctx, payload)
	if err != nil {
		return err
	}

	newRegisteredClient, err := c.service.Register(ctx.Request().Context(), payload)
	if err != nil {
		return err
	}

	var responseDto dto.RegisterSuccessResponse
	err = helper.Mapper(&responseDto, newRegisteredClient)
	if err != nil {
		return err
	}

	return response.Created(ctx, responseDto)
}

func (c *AuthController) Login(ctx echo.Context) error {
	payload := &dto.LoginRequest{}
	if err := helper.BindAndValidate(ctx, payload); err != nil {
		return err
	}

	clientSession, err := c.service.Login(ctx.Request().Context(), payload)
	if err != nil {
		return err
	}

	var responseDto dto.LoginSuccessResponse
	err = helper.Mapper(&responseDto, clientSession)
	if err != nil {
		return err
	}

	return response.Success(ctx, responseDto)
}
