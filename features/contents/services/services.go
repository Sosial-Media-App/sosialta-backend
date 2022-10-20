package services

import (
	"errors"

	"github.com/Sosial-Media-App/sosialta/config"
	"github.com/Sosial-Media-App/sosialta/features/contents/domain"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type contentServices struct {
	qry domain.Repository
}

func New(repo domain.Repository) domain.Services {
	return &contentServices{
		qry: repo,
	}
}

func (srv *contentServices) GetContent() ([]domain.Core, error) {
	res, err := srv.qry.Get()
	if err != nil {
		log.Error(err.Error())
		return nil, errors.New(config.NO_DATA)
	}

	return res, nil
}

func (srv *contentServices) GetContentDetail(id uint) (domain.Core, error) {
	res, err := srv.qry.GetDetail(id)
	if err != nil {
		return domain.Core{}, errors.New(config.NO_DATA)
	}
	return res, nil
}

func (srv *contentServices) AddContent(newContent domain.Core) (domain.Core, error) {
	res, err := srv.qry.Insert(newContent)

	if err != nil {
		return domain.Core{}, errors.New(config.DUPLICATED_DATA)
	}

	return res, nil
}

func (srv *contentServices) UpdateContent(updateData domain.Core) (domain.Core, error) {
	res, err := srv.qry.Update(updateData)

	if err != nil {
		return domain.Core{}, errors.New(config.DATABASE_ERROR)
	}

	return res, nil
}

func (srv *contentServices) DeleteContent(id uint) error {
	err := srv.qry.Delete(id)
	if err != nil {
		return errors.New(config.NO_DATA)
	}
	return nil
}

func (srv *contentServices) ExtractToken(c echo.Context) uint {
	token := c.Get("user").(*jwt.Token)
	if token.Valid {
		claim := token.Claims.(jwt.MapClaims)
		return uint(claim["id"].(float64))
	}

	return 0
}
