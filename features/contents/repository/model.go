package repository

import (
	"github.com/Sosial-Media-App/sosialta/features/contents/domain"
	"gorm.io/gorm"
)

type Content struct {
	gorm.Model
	IdUser       uint
	StoryType    string
	StoryDetail  string
	StoryPicture string
}

type Comment struct {
	gorm.Model
	IdUser    uint
	IdContent uint
	Comment   string
}

func FromDomain(dc domain.Core) Content {
	return Content{
		Model:        gorm.Model{ID: dc.ID},
		IdUser:       dc.ID,
		StoryType:    dc.StoryType,
		StoryDetail:  dc.StoryDetail,
		StoryPicture: dc.StoryPicture,
	}
}

func ToDomain(c Content) domain.Core {
	return domain.Core{
		ID:           c.ID,
		IdUser:       c.ID,
		StoryType:    c.StoryType,
		StoryDetail:  c.StoryDetail,
		StoryPicture: c.StoryPicture,
	}
}

func ToDomainArray(u Content, c []Comment) domain.Core {
	// var hasil domain.DetailCore
	var cContext []domain.CommentCore

	for _, val := range c {
		cContext = append(cContext, domain.CommentCore{
			ID: val.ID, IdUser: val.IdUser, IdContent: val.IdContent, Comment: val.Comment})
	}
	var res domain.Core = domain.Core{
		ID: u.ID, IdUser: u.IdUser, StoryType: u.StoryType, StoryDetail: u.StoryDetail, StoryPicture: u.StoryPicture, DetailCore: domain.DetailCore{cContext}}

	return res
}
