package repository

import (
	"github.com/Sosial-Media-App/sosialta/features/contents/domain"
	"gorm.io/gorm"
)

type repoQuery struct {
	db gorm.DB
}

func New(dbConn *gorm.DB) domain.Repository {
	return &repoQuery{
		db: *dbConn,
	}
}

func (rq *repoQuery) Insert(newContent domain.Core) (domain.Core, error) {

	return domain.Core{}, nil
}

func (rq *repoQuery) Update(updateData domain.Core, id uint) (domain.Core, error) {

	return domain.Core{}, nil
}

func (rq *repoQuery) Delete(id uint) error {

	return nil
}
func (rq *repoQuery) Get(newContent domain.Core) (domain.Core, error) {

	return domain.Core{}, nil
}

func (rq *repoQuery) GetDetail(newContent domain.Core) (domain.Core, error) {

	return domain.Core{}, nil
}
