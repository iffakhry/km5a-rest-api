package controller

import (
	"fakhry/clean-arch/features/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userUsecase user.UseCaseInterface
}

func New(userUC user.UseCaseInterface) *UserController {
	return &UserController{
		userUsecase: userUC,
	}
}

func (handler *UserController) CreateUser(c echo.Context) error {
	input := new(UserRequest)
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data",
		})
	}

	data := user.Core{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		Phone:    input.Phone,
	}
	err := handler.userUsecase.Create(data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "error insert data",
		})
	}
	return c.JSON(http.StatusCreated, map[string]any{
		"message": "success insert data",
	})
}
