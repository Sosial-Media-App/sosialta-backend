package delivery

import (
	"errors"
	"net/http"

	"github.com/Sosial-Media-App/sosialta/features/contents/domain"
	"github.com/labstack/echo/v4"
)

type contentHandler struct {
	srv domain.Services
}

func New(e *echo.Echo, srv domain.Services) {
	handler := contentHandler{srv: srv}
	e.GET("/Contents/:Contentname", handler.GetContent())
	// e.POST("/Contents", handler.RegiterContent())
	// e.POST("/login", handler.LoginContent())
	// e.PUT("/Contents", handler.UpdateDataContent(), middleware.JWT([]byte("Sosialta!!!12")))
	// e.DELETE("/Contents/:id", handler.DeactiveContent(), middleware.JWT([]byte("Sosialta!!!12")))
}

func (cs *contentHandler) RegiterContent() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusInternalServerError, errors.New("test"))
	}
}

func (cs *contentHandler) UpdateDataContent() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusInternalServerError, errors.New("test"))
	}
}

func (cs *contentHandler) GetContent() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusInternalServerError, errors.New("test"))
	}
}

func (cs *contentHandler) DeactiveContent() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusInternalServerError, errors.New("test"))
	}
}
