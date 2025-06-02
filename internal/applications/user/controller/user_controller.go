package controller

import (
	authsvc "mceasy/internal/applications/auth/service"
	"mceasy/internal/applications/user/dto"
	"mceasy/internal/applications/user/service"
	"mceasy/internal/helper"
	"mceasy/internal/helper/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	service     service.UserService
	authService authsvc.AuthService
}

func NewUserController(
	service service.UserService,
	authService authsvc.AuthService,
) *UserController {
	return &UserController{
		service:     service,
		authService: authService,
	}
}

// Create is controller to create new user.
//
//	@summary		Create new user
//	@description	Create new user
//	@tags			User
//	@accept			json
//	@produce		json
//	@param			X-Client-Key	header		string			true	"Client Key"
//	@param			body			body		dto.UserRequest	true	"Create User DTO"
//	@success		201				{object}	response.body{data=dto.UserResponse}
//	@failure		400				{object}	response.body
//	@failure		500				{object}	response.body
//	@router			/user [post]
func (c *UserController) Create(ctx echo.Context) error {
	request := new(dto.UserRequest)
	err := helper.BindAndValidate(ctx, request)
	if err != nil {
		return err
	}

	created, token, err := c.service.Create(ctx.Request().Context(), request)
	if err != nil {
		return response.Error(ctx, http.StatusBadRequest, err, nil)
	}

	var responseDto = new(dto.UserResponse)
	err = helper.Mapper(responseDto, created) // No need to use '&responseDto'
	if err != nil {
		return response.Error(ctx, http.StatusBadRequest, err, nil)
	}

	responseDto.Token = token

	return response.Created(ctx, responseDto)
}

// Update is controller to update a user.
//
//	@summary		Update a user
//	@description	Update a user
//	@tags			User
//	@accept			json
//	@produce		json
//	@param			X-Client-Key	header		string			true	"Client Key"
//	@param			id				path		int				true	"User's ID"
//	@param			body			body		dto.UserRequest	true	"Update User DTO"
//	@success		200				{object}	response.body{data=dto.UserUpdateResponse}
//	@failure		400				{object}	response.body
//	@failure		400				{object}	response.body
//	@failure		500				{object}	response.body
//	@router			/user/{id} [put]
func (c *UserController) Update(ctx echo.Context) error {
	request := new(dto.UserRequest)
	err := helper.BindAndValidate(ctx, request)
	if err != nil {
		return err
	}

	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return err
	}

	updated, err := c.service.Update(ctx.Request().Context(), uint64(id), request)
	if err != nil {
		return err
	}

	var responseDto = new(dto.UserUpdateResponse)
	err = helper.Mapper(&responseDto, updated)
	if err != nil {
		return err
	}

	return response.Success(ctx, responseDto)
}

// Delete is controller to delete a user.
//
//	@summary		Delete a user
//	@description	Delete a user
//	@tags			User
//	@accept			json
//	@produce		json
//	@param			X-Client-Key	header		string	true	"Client Key"
//	@param			id				path		int		true	"User's ID"
//	@success		200				{object}	response.body{data=dto.UserDeleteResponse}
//	@failure		400				{object}	response.body
//	@failure		500				{object}	response.body
//	@router			/user/{id} [delete]
func (c *UserController) Delete(ctx echo.Context) error {

	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return err
	}

	deleted, err := c.service.Delete(ctx.Request().Context(), uint64(id))
	if err != nil {
		return err
	}

	var responseDto = new(dto.UserDeleteResponse)
	err = helper.Mapper(&responseDto, deleted)
	if err != nil {
		return err
	}

	responseDto.IsDeleted = true

	return response.Success(ctx, responseDto)
}

// GetById is controller to get a user by its id.
//
//	@summary		Get a user by id
//	@description	Get a user by id
//	@tags			User
//	@accept			json
//	@produce		json
//	@param			X-Client-Key	header		string	true	"Client Key"
//	@param			id				path		int		true	"User's ID"
//	@success		200				{object}	response.body{data=dto.UserUpdateResponse}
//	@failure		400				{object}	response.body
//	@failure		500				{object}	response.body
//	@router			/user/{id} [get]
func (c *UserController) GetById(ctx echo.Context) error {

	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return err
	}

	result, err := c.service.GetById(ctx.Request().Context(), uint64(id))
	if err != nil {
		return err
	}

	return response.Success(ctx, result)
}

// GetAll is controller to get list of users.
//
//	@summary		Get all users
//	@description	Get all users
//	@tags			User
//	@accept			json
//	@produce		json
//	@param			X-Client-Key	header		string	true	"Client Key"
//	@success		200				{object}	response.body{data=[]dto.UserResponse}
//	@failure		400				{object}	response.body
//	@failure		500				{object}	response.body
//	@router			/user [get]
func (c *UserController) GetAll(ctx echo.Context) error {
	results, err := c.service.GetAll(ctx.Request().Context())
	if err != nil {
		return err
	}

	var responseDtos []*dto.UserResponse
	for _, result := range results {
		responseDto := new(dto.UserResponse)
		err = helper.Mapper(responseDto, result)
		if err != nil {
			return err
		}
		responseDtos = append(responseDtos, responseDto)
	}

	return response.Success(ctx, results)
}

// Login is controller to logging user in.
//
//	@summary		Login
//	@description	Login
//	@tags			User
//	@accept			json
//	@produce		json
//	@param			X-Client-Key	header		string					true	"Client Key"
//	@param			body			body		dto.UserLoginRequest	true	"Login User DTO"
//	@success		200				{object}	response.body{data=dto.UserResponse}
//	@failure		400				{object}	response.body
//	@failure		500				{object}	response.body
//	@router			/user/login [post]
func (c *UserController) Login(ctx echo.Context) error {
	request := new(dto.UserLoginRequest)
	err := helper.BindAndValidate(ctx, request)
	if err != nil {
		return err
	}

	user, token, err := c.service.Login(ctx.Request().Context(), request)
	if err != nil {
		return response.Error(ctx, http.StatusBadRequest, err, nil)
	}

	var responseDto = new(dto.UserLoginResponse)
	err = helper.Mapper(responseDto, user)
	if err != nil {
		return response.Error(ctx, http.StatusBadRequest, err, nil)
	}

	responseDto.Token = token

	return response.Success(ctx, responseDto)

}
