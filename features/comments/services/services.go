package services

import (
	"errors"
	"strings"

	"github.com/Sosial-Media-App/sosialta/config"
	"github.com/Sosial-Media-App/sosialta/features/comments/domain"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type commentService struct {
	qry domain.Repository
}

func New(repo domain.Repository) domain.Services {
	return &commentService{
		qry: repo,
	}
}

func (cs *commentService) GetComment(id_content uint) ([]domain.Core, error) {
	res, err := cs.qry.Get(id_content)
	if err != nil {
		log.Error(err.Error())
		if strings.Contains(err.Error(), "table") {
			return nil, errors.New("database error")
		} else if strings.Contains(err.Error(), "found") {
			return nil, errors.New("no data")
		}
	}

	return res, nil
}

func (cs *commentService) AddComment(newComment domain.Core) (domain.Core, error) {
	res, err := cs.qry.Insert(newComment)

	if err != nil {
		return domain.Core{}, errors.New(config.DUPLICATED_DATA)
	}

	return res, nil
}

func (cs *commentService) UpdateComment(updateData domain.Core, id uint) (domain.Core, error) {
	res, err := cs.qry.Update(updateData, id)
	if err != nil {
		if strings.Contains(err.Error(), config.DUPLICATED_DATA) {
			return domain.Core{}, errors.New(config.REJECTED_DATA)
		}
	}

	return res, nil
}

func (cs *commentService) DeleteComment(id uint) error {
	err := cs.qry.Delete(id)
	if err != nil {
		return errors.New("data not found")
	}
	return nil
}

func (cs *commentService) ExtractToken(c echo.Context) uint {
	token := c.Get("user").(*jwt.Token)
	if token.Valid {
		claim := token.Claims.(jwt.MapClaims)
		return uint(claim["id"].(float64))
	}

	return 0
}
