package delivery

import (
	"context"
	"errors"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/Sosial-Media-App/sosialta/features/contents/domain"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type contentHandler struct {
	srv domain.Services
}

func New(e *echo.Echo, srv domain.Services) {
	handler := contentHandler{srv: srv}
	e.GET("/contents", handler.GetContent())
	e.POST("/contents", handler.RegiterContent(), middleware.JWT([]byte("Sosialta!!!12")))
	e.GET("/contents/:id", handler.GetContentDetail())
	e.PUT("/contents/:id", handler.UpdateDataContent(), middleware.JWT([]byte("Sosialta!!!12")))
	e.DELETE("/contents/:id", handler.DeactiveContent(), middleware.JWT([]byte("Sosialta!!!12")))
}

func (cs *contentHandler) RegiterContent() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input RegiterFormat
		input.StoryType = c.FormValue("story_type")
		input.StoryDetail = c.FormValue("story_detail")
		input.IdUser = cs.srv.ExtractToken(c)
		file, err := c.FormFile("file")
		if err != nil {
			return err
		}

		src, err := file.Open()
		if err != nil {
			return err
		} else {
			input.StoryPicture = "https://sosialtabucket.s3.ap-southeast-1.amazonaws.com/myfiles/" + strings.ReplaceAll(file.Filename, " ", "+")
		}
		defer src.Close()

		s3Config := &aws.Config{
			Region:      aws.String(os.Getenv("AWS_REGION")),
			Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_USER"), os.Getenv("AWS_KEY"), ""),
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

		cnv := ToDomain(input)
		res, err := cs.srv.AddContent(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailedResponse(err.Error()))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("berhasil register", ToResponse(res, "register")))
	}
}

func (cs *contentHandler) UpdateDataContent() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input UpdateFormat
		input.StoryDetail = c.FormValue("story_detail")
		ID := c.Param("id")
		u64, err := strconv.ParseUint(ID, 10, 32)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailedResponse("Wrong id content"))
		}
		input.ID = uint(u64)
		file, err := c.FormFile("file")
		if err != nil {
			return err
		}

		src, err := file.Open()
		if err != nil {
			return err
		} else {
			input.StoryPicture = "https://sosialtabucket.s3.ap-southeast-1.amazonaws.com/myfiles/" + strings.ReplaceAll(file.Filename, " ", "+")
		}
		defer src.Close()

		s3Config := &aws.Config{
			Region:      aws.String(os.Getenv("AWS_REGION")),
			Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_USER"), os.Getenv("AWS_KEY"), ""),
		}

		s3Session := session.New(s3Config)

		uploader := s3manager.NewUploader(s3Session)
		inputData := &s3manager.UploadInput{
			Bucket: aws.String("sosialtabucket"),           // bucket's name
			Key:    aws.String("myfiles/" + file.Filename), // files destination location
			Body:   src,                                    // content of the file

		}
		_, _ = uploader.UploadWithContext(context.Background(), inputData)

		cnv := ToDomain(input)
		res, err := cs.srv.UpdateContent(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailedResponse(err.Error()))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("berhasil update", ToResponse(res, "update")))
	}
}

func (cs *contentHandler) GetContent() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := cs.srv.GetContent()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, errors.New("test"))
		}
		return c.JSON(http.StatusCreated, SuccessResponse("Success get data", ToResponseContent(res, "all")))
	}
}

func (cs *contentHandler) GetContentDetail() echo.HandlerFunc {
	return func(c echo.Context) error {
		cnv, errCnv := strconv.Atoi(c.Param("id"))
		if errCnv != nil {
			return c.JSON(http.StatusInternalServerError, "cant convert id")
		}

		res, err := cs.srv.GetContentDetail(uint(cnv))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, errors.New("test"))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("Success get data", ToResponse(res, "getdetail")))
	}
}

func (cs *contentHandler) DeactiveContent() echo.HandlerFunc {
	return func(c echo.Context) error {
		cnv, errCnv := strconv.Atoi(c.Param("id"))
		if errCnv != nil {
			return c.JSON(http.StatusInternalServerError, "cant convert id")
		}

		err := cs.srv.DeleteContent(uint(cnv))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success delete data.",
		})
	}
}
