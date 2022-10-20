package delivery

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

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

func (cs *contentHandler) GetContent() echo.HandlerFunc {
	return func(c echo.Context) error {
		page, err := strconv.Atoi(c.QueryParam("page"))
		if err != nil {
			page = 0
		}
		res, err := cs.srv.GetContent(page)
		log.Println(res, err)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, errors.New("test"))
		}
		return c.JSON(http.StatusCreated, SuccessResponse("Success get data", ToResponseContent(res, "all")))
	}
}

func (cs *contentHandler) RegiterContent() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input RegiterFormat
		input.StoryType = c.FormValue("story_type")
		input.StoryDetail = c.FormValue("story_detail")
		input.IdUser = cs.srv.ExtractToken(c)
		file, err := c.FormFile("story_picture")
		if err == nil {
			src, err := file.Open()
			if err != nil {
				input.StoryPicture = "https://sosialtabucket.s3.ap-southeast-1.amazonaws.com/myfiles/casting-couch.jpg"
			}
			defer src.Close()

			s3Config := &aws.Config{
				Region:      aws.String(os.Getenv("AWS_REGION")),
				Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_USER"), os.Getenv("AWS_KEY"), ""),
			}
			temp := time.Now().Format("02 Jan 06 15:04")
			input.StoryPicture = "https://sosialtabucket.s3.ap-southeast-1.amazonaws.com/myfiles/" + temp + strings.ReplaceAll(file.Filename, " ", "+")
			s3Session := session.New(s3Config)

			uploader := s3manager.NewUploader(s3Session)
			inputData := &s3manager.UploadInput{
				Bucket: aws.String("sosialtabucket"),                  // bucket's name
				Key:    aws.String("myfiles/" + temp + file.Filename), // files destination location
				Body:   src,                                           // content of the file

			}
			_, _ = uploader.UploadWithContext(context.Background(), inputData)
		}
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
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailedResponse("Wrong page number"))
		}
		input.ID = uint(id)
		file, err := c.FormFile("story_picture")
		if err == nil {
			src, err := file.Open()
			if err != nil {
				input.StoryPicture = "https://sosialtabucket.s3.ap-southeast-1.amazonaws.com/myfiles/casting-couch.jpg"
			}
			defer src.Close()

			s3Config := &aws.Config{
				Region:      aws.String(os.Getenv("AWS_REGION")),
				Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_USER"), os.Getenv("AWS_KEY"), ""),
			}

			temp := time.Now().Format("02 Jan 06 15:04")
			input.StoryPicture = "https://sosialtabucket.s3.ap-southeast-1.amazonaws.com/myfiles/" + temp + strings.ReplaceAll(file.Filename, " ", "+")
			s3Session := session.New(s3Config)

			uploader := s3manager.NewUploader(s3Session)
			inputData := &s3manager.UploadInput{
				Bucket: aws.String("sosialtabucket"),                  // bucket's name
				Key:    aws.String("myfiles/" + temp + file.Filename), // files destination location
				Body:   src,                                           // content of the file

			}
			_, _ = uploader.UploadWithContext(context.Background(), inputData)
		}

		cnv := ToDomain(input)
		res, err := cs.srv.UpdateContent(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailedResponse(err.Error()))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("berhasil update", ToResponse(res, "update")))
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
