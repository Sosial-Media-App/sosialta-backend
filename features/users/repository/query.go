package repository

import (
	"github.com/Sosial-Media-App/sosialta/features/users/domain"
	"gorm.io/gorm"
)

type repoQuery struct {
	db *gorm.DB
}

func New(dbConn *gorm.DB) domain.Repository {
	return &repoQuery{
		db: dbConn,
	}
}

func (rq *repoQuery) Insert(newUser domain.Core) (domain.Core, error) {

	return newUser, nil
}

func (rq *repoQuery) Login(newUser domain.Core) (domain.Core, string, error) {

	return newUser, "token", nil
}

func (rq *repoQuery) Update(newUser domain.Core) (domain.Core, error) {

	return newUser, nil
}

func (rq *repoQuery) Delete(newUser domain.Core) (domain.Core, error) {

	return newUser, nil
}
