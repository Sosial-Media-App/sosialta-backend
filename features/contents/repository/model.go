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

func ToDomainArray(u []Content, c []Comment) []domain.Core {
	var res []domain.Core
	var cComment []domain.CommentCore
	var j, i int = 3, 0
	for _, val := range u {
		for i < j && i < len(c) {
			if c[i].IdContent == val.ID {
				cComment = append(cComment, domain.CommentCore{
					ID: c[i].ID, IdUser: c[i].IdUser, IdContent: c[i].IdContent, Comment: c[i].Comment})
			}
			i++
		}
		res = append(res, domain.Core{ID: val.ID, IdUser: val.IdUser, StoryType: val.StoryType, StoryDetail: val.StoryDetail,
			StoryPicture: val.StoryPicture, DetailCore: domain.DetailCore{cComment}})
		j += 3
	}
	// var res domain.Core = domain.Core{
	// 	ID: u.ID, IdUser: u.IdUser, StoryType: u.StoryType, StoryDetail: u.StoryDetail, StoryPicture: u.StoryPicture, DetailCore: domain.DetailCore{cContext}}

	return res
}
