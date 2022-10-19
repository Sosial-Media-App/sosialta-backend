package delivery

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/Sosial-Media-App/sosialta/config"
	"github.com/Sosial-Media-App/sosialta/features/users/domain"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type userHandler struct {
	srv domain.Services
}

func New(e *echo.Echo, srv domain.Services) {
	handler := userHandler{srv: srv}
	e.GET("/users", handler.GetUser())
	e.GET("/users/:username", handler.GetUser())
	e.POST("/users", handler.RegiterUser())
	e.POST("/login", handler.LoginUser())
	e.PUT("/users", handler.UpdateDataUser(), middleware.JWT([]byte("Sosialta!!!12")))
	e.DELETE("/users/:id", handler.DeactiveUser(), middleware.JWT([]byte("Sosialta!!!12")))
}

func (us *userHandler) GetUser() echo.HandlerFunc {
	//mendapatkan detail profile
	return func(c echo.Context) error {
		var resQry RegiterFormat
		var myUser bool = false
		resQry.Username = c.QueryParam("username")
		cnv := ToDomain(resQry)
		res, err := us.srv.GetUser(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailedResponse(err.Error()))
		}

		// id := us.srv.ExtractToken(c)
		// if id == res.ID {
		// 	//memberi pembembeda sementara antara user dan user lain
		// 	myUser = true
		// }

		return c.JSON(http.StatusCreated, SuccessResponse("berhasil register", ToResponseUser(res, myUser, "getuser")))
	}
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
			return c.JSON(http.StatusBadRequest, config.PARSE_DATA)
		}

		input.UserPicture = "https://sosialtabucket.s3.ap-southeast-1.amazonaws.com/myfiles/Screenshot+(316).png"
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
		var input UpdateFormat
		input.Fullname = c.FormValue("fullname")
		input.Username = c.FormValue("username")
		input.Email = c.FormValue("email")
		input.Password = c.FormValue("password")
		input.Phone = c.FormValue("phone")
		input.Dob = c.FormValue("dob")

		file, err := c.FormFile("file")
		if err != nil {
			return err
		}

		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		s3Config := &aws.Config{
			Region:      aws.String("ap-southeast-1"),
			Credentials: credentials.NewStaticCredentials("AKIAUFHWMWYWKGW2OIUP", "WInFzSVwxTiaOmoOoLyQ7jtk0nAkuH9WNQc9zJDM", ""),
		}

		s3Session := session.New(s3Config)

		uploader := s3manager.NewUploader(s3Session)
		inputData := &s3manager.UploadInput{
			Bucket: aws.String("sosialtabucket"),           // bucket's name
			Key:    aws.String("myfiles/" + file.Filename), // files destination location
			Body:   src,                                    // content of the file
			// ACL:    aws.String("Objects - List"),
			// ContentType: aws.String("image/png"), // content type
		}
		_, _ = uploader.UploadWithContext(context.Background(), inputData)
		input.UserPicture = "https://sosialtabucket.s3.ap-southeast-1.amazonaws.com/myfiles/" + strings.ReplaceAll(file.Filename, " ", "+")
		cnv := ToDomain(input)

		userId := us.srv.ExtractToken(c)
		res, err := us.srv.UpdateUser(cnv, userId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, ToResponse(res, "update"))
	}
}

func (us *userHandler) DeactiveUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		cnv, errCnv := strconv.Atoi(c.Param("id"))
		if errCnv != nil {
			return c.JSON(http.StatusInternalServerError, "cant convert id")
		}

		err := us.srv.DeleteUser(uint(cnv))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success delete data.",
		})
	}
}
