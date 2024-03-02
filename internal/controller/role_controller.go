package controller

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"management-api/internal/model"
	"management-api/internal/service"
)

type RoleController struct {
	service service.IRoleService
}

func NewRoleController(service service.IRoleService) *RoleController {
	return &RoleController{service: service}
}

// GetAllRole godoc
//
//	@Summary		Get all roles
//	@Description	Get all roles
//	@Tags			Role
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	model.Role
//	@Router			/role [get]
func (controller *RoleController) GetAllRole(c echo.Context) error {
	res, err := controller.service.GetAllRole()
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, res)
}

// GetRoleById godoc
//
//	@Summary		Get roles by ID
//	@Description	Get roles by ID
//	@Tags			Role
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	[]model.Role
//	@Param			id	path	string	true	"role ID"
//	@Router			/role/id/:id [get]
func (controller *RoleController) GetRoleById(c echo.Context) error {
	id := c.Param("id")

	res, err := controller.service.GetRoleById(id)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, res)
}

// UpsertRole godoc
//
//	@Summary		Upsert role
//	@Description	Upsert role
//	@Tags			Role
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.Role
//	@Router			/role [post]
func (controller *RoleController) UpsertRole(c echo.Context) error {
	var requestBody model.Role
	err := json.NewDecoder(c.Request().Body).Decode(&requestBody)
	if err != nil {
		return c.JSON(400, err)
	}
	err = controller.service.UpsertRole(&requestBody)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, requestBody)
}

// DeleteRole godoc
//
//	@Summary		Delete role
//	@Description	Delete role
//	@Tags			Role
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	string
//	@Param			id	path		string	true	"Role ID"
//	@Router			/role/id/:id [delete]
func (controller *RoleController) DeleteRole(c echo.Context) error {
	id := c.Param("id")

	err := controller.service.DeleteRole(id)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, "Deleted")
}
