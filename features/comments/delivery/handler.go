package delivery

import (
	"net/http"
	"strconv"

	"github.com/Sosial-Media-App/sosialta/config"
	"github.com/Sosial-Media-App/sosialta/features/comments/domain"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type commentHandler struct {
	srv domain.Services
}

func New(e *echo.Echo, srv domain.Services) {
	handler := commentHandler{
		srv: srv,
	}

	e.GET("/comments/:id", handler.GetComment())
	e.POST("/comments", handler.AddComment(), middleware.JWT([]byte("Sosialta!!!12")))
	e.PUT("/comments", handler.UpdateComment(), middleware.JWT([]byte("Sosialta!!!12")))
	e.DELETE("/comments/:id", handler.DeactiveComment(), middleware.JWT([]byte("Sosialta!!!12")))
}

func (cs *commentHandler) GetComment() echo.HandlerFunc {
	return func(c echo.Context) error {
		cnv, errCnv := strconv.Atoi(c.Param("id"))
		if errCnv != nil {
			return c.JSON(http.StatusInternalServerError, "cant convert id")
		}

		res, err := cs.srv.GetComment(uint(cnv))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailedResponse(err.Error()))
		}

		return c.JSON(http.StatusOK, SuccessResponse("berhasil get comment", ToResponseComment(res, "get")))
	}
}

func (cs *commentHandler) AddComment() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input RegiterFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, config.PARSE_DATA)
		}

		input.IdUser = cs.srv.ExtractToken(c)

		cnv := ToDomain(input)
		res, err := cs.srv.AddComment(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailedResponse(err.Error()))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("berhasil add comment", ToResponse(res, "register")))
	}
}

func (cs *commentHandler) UpdateComment() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input UpdateFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, config.PARSE_DATA)
		}

		input.IdUser = cs.srv.ExtractToken(c)

		cnv := ToDomain(input)
		res, err := cs.srv.UpdateComment(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailedResponse(err.Error()))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("berhasil update comment", ToResponse(res, "update")))
	}
}

func (cs *commentHandler) DeactiveComment() echo.HandlerFunc {
	return func(c echo.Context) error {
		cnv, errCnv := strconv.Atoi(c.Param("id"))
		if errCnv != nil {
			return c.JSON(http.StatusInternalServerError, "cant convert id")
		}

		err := cs.srv.DeleteComment(uint(cnv))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success delete data.",
		})
	}
}
