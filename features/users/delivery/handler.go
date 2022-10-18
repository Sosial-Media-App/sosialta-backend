package delivery

import (
	"net/http"

	"github.com/Sosial-Media-App/sosialta/config"
	"github.com/Sosial-Media-App/sosialta/features/users/domain"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type userHandler struct {
	srv domain.Services
}

func New(e *echo.Echo, srv domain.Services) {
	handler := userHandler{srv: srv}
	e.POST("/users", handler.RegiterUser())
	e.POST("/login", handler.LoginUser())
	e.PUT("/users", handler.UpdateDataUser(), middleware.JWT([]byte("Sosialta!!!12")))
}

func (us *userHandler) LoginUser() echo.HandlerFunc {
	//autentikasi user login
	return func(c echo.Context) error {
		var resQry RegiterFormat
		if err := c.Bind(&resQry); err != nil {
			return c.JSON(http.StatusBadRequest, FailedResponse("cannot bind input"))
		}

		cnv := ToDomain(resQry)
		res, err := us.srv.Login(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailedResponse(err.Error()))
		}
		token := us.srv.GenerateToken(res.ID, res.Username)

		return c.JSON(http.StatusCreated, SuccessResponse("berhasil register", ToResponseLogin(res, token, "login")))
	}
}

func (us *userHandler) RegiterUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input RegiterFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailedResponse("cannot bind input"))
		}
		cnv := ToDomain(input)
		res, err := us.srv.AddUser(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailedResponse(err.Error()))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("berhasil register", ToResponse(res, "register")))
	}
}

func (us *userHandler) UpdateDataUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var updateData UpdateFormat
		if err := c.Bind(&updateData); err != nil {
			return c.JSON(http.StatusBadRequest, config.PARSE_DATA)
		}
		cnv := ToDomain(updateData)

		userId := us.srv.ExtractToken(c)
		res, err := us.srv.UpdateUser(cnv, userId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success update data.",
			"data":    res,
		})
	}
}
