package repository

import (
	"github.com/Sosial-Media-App/sosialta/features/contents/domain"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string
	Email       string
	Password    string
	Fullname    string
	Phone       string
	Dob         string
	UserPicture string
}
type Content struct {
	gorm.Model
	IdUser       uint
	Username     string
	StoryType    string
	StoryDetail  string
	StoryPicture string
}

type Comment struct {
	gorm.Model
	IdUser    uint
	Username  string
	IdContent uint
	Comment   string
}

func FromDomain(dc domain.Core) Content {
	return Content{
		Model:        gorm.Model{ID: dc.ID},
		IdUser:       dc.IdUser,
		Username:     dc.Username,
		StoryType:    dc.StoryType,
		StoryDetail:  dc.StoryDetail,
		StoryPicture: dc.StoryPicture,
	}
}

func ToDomain(c Content) domain.Core {
	return domain.Core{
		ID:           c.ID,
		IdUser:       c.IdUser,
		Username:     c.Username,
		StoryType:    c.StoryType,
		StoryDetail:  c.StoryDetail,
		StoryPicture: c.StoryPicture,
	}
}

func ToDomainArray(u []Content, c []Comment) []domain.Core {
	var res []domain.Core
	var cComment []domain.CommentCore
	var j int = 0
	for _, val := range u {
		for i := 0; i < len(c) && j < 3; i++ {
			if c[i].IdContent == val.ID {
				cComment = append(cComment, domain.CommentCore{
					ID: c[i].ID, IdUser: c[i].IdUser, IdContent: c[i].IdContent, Username: c[i].Username, Comment: c[i].Comment})
			}
		}
		res = append(res, domain.Core{ID: val.ID, IdUser: val.IdUser, Username: val.Username, StoryType: val.StoryType, StoryDetail: val.StoryDetail,
			StoryPicture: val.StoryPicture, DetailCore: domain.DetailCore{cComment}})
	}
	return res
}

func ToDomainDetail(u Content, c []Comment) domain.Core {
	// var hasil domain.DetailCore
	var cComment []domain.CommentCore

	for _, val := range c {
		cComment = append(cComment, domain.CommentCore{
			ID: val.ID, IdUser: val.ID, Username: val.Username, IdContent: val.IdContent, Comment: val.Comment})
	}
	var res domain.Core = domain.Core{
		IdUser:       u.ID,
		Username:     u.Username,
		StoryType:    u.StoryType,
		StoryDetail:  u.StoryDetail,
		StoryPicture: u.StoryPicture, DetailCore: domain.DetailCore{cComment}}

	return res
}
