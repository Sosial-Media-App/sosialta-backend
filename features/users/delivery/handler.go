package delivery

import (
	"errors"
	"net/http"

	"github.com/Sosial-Media-App/sosialta/config"
	"github.com/Sosial-Media-App/sosialta/features/users/domain"
	"github.com/labstack/echo/v4"
)

type userHandler struct {
	srv domain.Services
}

func New(e *echo.Echo, srv domain.Services) {
	handler := userHandler{srv: srv}
	e.POST("/users", handler.RegiterUser())
	e.PUT("/users", handler.UpdateDataUser())
}

func (us *userHandler) RegiterUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusBadRequest, errors.New("fail"))
	}
}

func (us *userHandler) LoginUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusBadRequest, errors.New("fail"))
	}
}

func (us *userHandler) UpdateDataUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var updateData domain.Core
		if err := c.Bind(&updateData); err != nil {
			return c.JSON(http.StatusBadRequest, config.PARSE_DATA)
		}

		res, err := us.srv.UpdateUser(updateData)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success update data.",
			"data":    res,
		})
	}
}
