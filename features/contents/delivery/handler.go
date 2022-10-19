package delivery

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/Sosial-Media-App/sosialta/features/contents/domain"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/labstack/echo/v4"
)

type contentHandler struct {
	srv domain.Services
}

func New(e *echo.Echo, srv domain.Services) {
	handler := contentHandler{srv: srv}
	e.GET("/Contents/:Contentname", handler.GetContent())
	e.POST("/Contents", handler.RegiterContent())
	// e.POST("/login", handler.LoginContent())
	// e.PUT("/Contents", handler.UpdateDataContent(), middleware.JWT([]byte("Sosialta!!!12")))
	// e.DELETE("/Contents/:id", handler.DeactiveContent(), middleware.JWT([]byte("Sosialta!!!12")))
}

func (cs *contentHandler) RegiterContent() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input RegiterFormat
		input.StoryType = c.FormValue("story_type")
		input.StoryDetail = c.FormValue("story_detail")

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
