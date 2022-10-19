package services

import (
	"errors"

	"github.com/Sosial-Media-App/sosialta/config"
	"github.com/Sosial-Media-App/sosialta/features/contents/domain"
	"github.com/labstack/echo/v4"
)

type contentServices struct {
	qry domain.Repository
}

func New(repo domain.Repository) domain.Services {
	return &contentServices{
		qry: repo,
	}
}

func (srv *contentServices) GetContent(newContent domain.Core) (domain.Core, error) {

	return domain.Core{}, nil
}

func (srv *contentServices) GetContentDetail(newContent domain.Core) (domain.Core, error) {

	return domain.Core{}, nil
}

func (srv *contentServices) AddContent(newContent domain.Core) (domain.Core, error) {
	res, err := srv.qry.Insert(newContent)

	if err != nil {
		return domain.Core{}, errors.New(config.DUPLICATED_DATA)
	}

	return res, nil
}

func (srv *contentServices) UpdateContent(updateData domain.Core, id uint) (domain.Core, error) {

	return domain.Core{}, nil
}

func (srv *contentServices) DeleteContent(id uint) error {

	return nil
}

func (srv *contentServices) ExtractToken(c echo.Context) uint {

	return 0
}
