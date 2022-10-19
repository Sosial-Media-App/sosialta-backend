package repository

import (
	"github.com/Sosial-Media-App/sosialta/features/contents/domain"
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

func (rq *repoQuery) Insert(newContent domain.Core) (domain.Core, error) {
	var cnv Content = FromDomain(newContent)
	if err := rq.db.Create(&cnv).Error; err != nil {
		return domain.Core{}, err
	}
	// selesai dari DB
	newContent = ToDomain(cnv)
	return newContent, nil
}

func (rq *repoQuery) Update(updateData domain.Core, id uint) (domain.Core, error) {

	return domain.Core{}, nil
}

func (rq *repoQuery) Delete(id uint) error {

	return nil
}
func (rq *repoQuery) Get() ([]domain.Core, error) {
	var resQry []Content
	var resQryComment []Comment
	if err := rq.db.Limit(20).Order("created_at desc").Find(&resQry).Error; err != nil {
		return nil, err
	}
	// selesai dari DB
	for _, val := range resQry {
		if err := rq.db.Limit(3).Order("created_at desc").Find(&resQryComment, "id_content = ?", val.ID).Error; err != nil {
			return nil, err
		}
	}
	res := ToDomainArray(resQry, resQryComment)
	return res, nil
}

func (rq *repoQuery) GetDetail(newContent domain.Core) (domain.Core, error) {

	return domain.Core{}, nil
}