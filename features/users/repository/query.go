package repository

import (
	"errors"

	"github.com/Sosial-Media-App/sosialta/config"
	"github.com/Sosial-Media-App/sosialta/features/users/domain"
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

func (rq *repoQuery) Get(newUser domain.Core) (domain.Core, error) {
	var resQry User = FromDomain(newUser)
	var resQryContent []Content
	if err := rq.db.First(&resQry, "username = ?", resQry.Username).Error; err != nil {
		return domain.Core{}, err
	}
	// selesai dari DB
	if err := rq.db.Find(&resQryContent, "id_user = ?", resQry.ID).Error; err != nil {
		return domain.Core{}, err
	}
	res := ToDomainArray(resQry, resQryContent)
	return res, nil
}

func (rq *repoQuery) Insert(newUser domain.Core) (domain.Core, error) {
	var cnv User = FromDomain(newUser)
	if err := rq.db.Create(&cnv).Error; err != nil {
		return domain.Core{}, err
	}
	// selesai dari DB
	newUser = ToDomain(cnv)
	return newUser, nil
}

func (rq *repoQuery) Login(newUser domain.Core) (domain.Core, error) {
	var resQry User
	if err := rq.db.First(&resQry, "username = ? OR email=?", newUser.Username, newUser.Email).Error; err != nil {
		return domain.Core{}, err
	}

	// selesai dari DB
	res := ToDomain(resQry)
	return res, nil
}

func (rq *repoQuery) Update(updateData domain.Core, id uint) (domain.Core, error) {
	var resQry User
	var resCty Content
	resQry = FromDomain(updateData)

	err := rq.db.Where("id = ?", id).Updates(resQry).Error
	if err != nil {
		log.Error(config.DATABASE_ERROR)
		return domain.Core{}, err
	}
	if resQry.Username != "" {
		rq.db.Where("username=?", resQry.Username).Updates(&resCty)
	}

	updateData = ToDomain(resQry)

	return updateData, nil
}

func (rq *repoQuery) Delete(id uint) error {
	err := rq.db.Where("id = ?", id).Delete(&User{})
	if err.Error != nil {
		return errors.New("cant delete data")
	}

	if err.RowsAffected < 1 {
		return errors.New("row not affected")
	}

	return nil
}
