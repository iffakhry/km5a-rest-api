package handler

import (
	"fakhry/clean-arch/features/user"
	"net/http"
	"strings"

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
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": err.Error(),
			})

		}
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "error insert data. " + err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, map[string]any{
		"message": "success insert data",
	})
}
