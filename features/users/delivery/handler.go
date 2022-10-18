package delivery

import (
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
	// e.PUT("/users", handler.UpdateDataUser())
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
		// token := us.srv.GenerateToken(res.ID)
		return c.JSON(http.StatusCreated, SuccessResponse("berhasil register", ToResponse(res, "login")))
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

// func (us *userHandler) UpdateDataUser() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		var updateData UpdateUserFormat
// 		if err := c.Bind(&updateData); err != nil {
// 			return c.JSON(http.StatusBadRequest, config.PARSE_DATA)
// 		}

// 		res, err := us.srv.UpdateUser(updateData)
// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, err.Error())
// 		}

// 		return c.JSON(http.StatusOK, map[string]interface{}{
// 			"message": "Success update data.",
// 			"data":    res,
// 		})
// 	}
// }
