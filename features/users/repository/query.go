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
	if err := rq.db.First(&resQry, "id = ?", newUser.ID).Error; err != nil {
		return domain.Core{}, err
	}

	// selesai dari DB
	res := ToDomain(resQry)
	return res, nil
}

func (rq *repoQuery) Update(updateData domain.Core) (domain.Core, error) {
	qryData := FromDomain(updateData)

	err := rq.db.Where("id = ?", qryData.ID).Updates(updateData).Error
	if err != nil {
		return domain.Core{}, err
	}

	updateData = ToDomain(qryData)

	return updateData, nil
}

func (rq *repoQuery) Delete(newUser domain.Core) (domain.Core, error) {

	return newUser, nil
}
