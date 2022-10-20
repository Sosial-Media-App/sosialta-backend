package repository

import (
	"github.com/Sosial-Media-App/sosialta/features/comments/domain"
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
type Comment struct {
	gorm.Model
	IdUser    uint
	Username  string
	IdContent uint
	Comment   string
}

func FromDomain(dc domain.Core) Comment {
	return Comment{
		Model:     gorm.Model{ID: dc.ID},
		IdUser:    dc.IdUser,
		Username:  dc.Username,
		IdContent: dc.IdContent,
		Comment:   dc.Comment,
	}
}

func ToDomain(c Comment) domain.Core {
	return domain.Core{
		ID:        c.ID,
		IdUser:    c.IdUser,
		Username:  c.Username,
		IdContent: c.IdContent,
		Comment:   c.Comment,
	}
}

func ToDomainArray(ac []Comment) []domain.Core {
	var res []domain.Core
	for _, val := range ac {
		res = append(res, domain.Core{ID: val.ID, IdUser: val.IdUser, Username: val.Username, IdContent: val.IdContent, Comment: val.Comment})
	}

	return res
}
