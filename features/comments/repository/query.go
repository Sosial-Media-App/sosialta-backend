package repository

import (
	"errors"

	"github.com/Sosial-Media-App/sosialta/config"
	"github.com/Sosial-Media-App/sosialta/features/comments/domain"
	"github.com/labstack/gommon/log"
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

func (rq *repoQuery) Get(id_content uint) ([]domain.Core, error) {
	var resQry []Comment
	if err := rq.db.Find(&resQry, "id_content = ?", id_content).Error; err != nil {
		return nil, err
	}
	// selesai dari DB
	res := ToDomainArray(resQry)
	return res, nil
}

func (rq *repoQuery) Insert(newComment domain.Core) (domain.Core, error) {
	var cnv Comment = FromDomain(newComment)
	if err := rq.db.Create(&cnv).Error; err != nil {
		return domain.Core{}, err
	}
	// selesai dari DB
	newComment = ToDomain(cnv)
	return newComment, nil
}

func (rq *repoQuery) Update(newComment domain.Core) (domain.Core, error) {
	var cnv Comment = FromDomain(newComment)
	err := rq.db.Where("id = ?", cnv.ID).Updates(cnv).Error
	if err != nil {
		log.Error(config.DATABASE_ERROR)
		return domain.Core{}, err
	}

	newComment = ToDomain(cnv)

	return newComment, nil
}

func (rq *repoQuery) Delete(id uint) error {
	err := rq.db.Where("id = ?", id).Delete(&Comment{})
	if err.Error != nil {
		return errors.New("cant delete data")
	}

	if err.RowsAffected < 1 {
		return errors.New("row not affected")
	}

	return nil
}
