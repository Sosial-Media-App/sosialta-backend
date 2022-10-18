package delivery

import (
	"errors"
	"net/http"

	"github.com/Sosial-Media-App/sosialta/features/users/domain"
	"github.com/labstack/echo/v4"
)

type userHandler struct {
	srv domain.Services
}

func New(e *echo.Echo, srv domain.Services) {
	handler := userHandler{srv: srv}
	e.POST("/users", handler.RegiterUser())
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
