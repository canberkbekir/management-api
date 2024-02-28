package controller

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"management-api/internal/model"
	"management-api/internal/service"
)

type UserController struct {
	service service.IUserService
}

func NewUserController(service service.IUserService) *UserController {
	return &UserController{service: service}
}

// GetAllUser godoc
//
//	@Summary		Get all Users
//	@Description	Get all Users
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	model.User
//	@Router			/user [get]
func (controller *UserController) GetAllUser(c echo.Context) error {
	res, err := controller.service.GetAllUser()
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, res)
}

// GetUserById godoc
//
//	@Summary		Get users by ID
//	@Description	Get users by ID
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	[]model.User
//	@Param			id	path	string	true	"User ID"
//	@Router			/user/id/:id [get]
func (controller *UserController) GetUserById(c echo.Context) error {
	id := c.Param("id")

	res, err := controller.service.GetUserById(id)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, res)
}

func (controller *UserController) UpsertUser(c echo.Context) error {
	var requestBody model.User
	err := json.NewDecoder(c.Request().Body).Decode(&requestBody)
	if err != nil {
		return c.JSON(400, err)
	}
	err = controller.service.UpsertUser(&requestBody)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, requestBody)
}

func (controller *UserController) DeleteUser(c echo.Context) error {
	id := c.Param("id")

	err := controller.service.DeleteUser(id)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, "Deleted")
}
