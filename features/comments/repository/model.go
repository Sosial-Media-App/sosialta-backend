package repository

import (
	"github.com/Sosial-Media-App/sosialta/features/comments/domain"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	IdUser    uint
	IdContent uint
	Comment   string
}

func FromDomain(dc domain.Core) Comment {
	return Comment{
		Model:     gorm.Model{ID: dc.ID},
		IdUser:    dc.IdUser,
		IdContent: dc.IdContent,
		Comment:   dc.Comment,
	}
}

func ToDomain(c Comment) domain.Core {
	return domain.Core{
		ID:        c.ID,
		IdUser:    c.IdUser,
		IdContent: c.IdContent,
		Comment:   c.Comment,
	}
}

func ToDomainArray(ac []Comment) []domain.Core {
	var res []domain.Core
	for _, val := range ac {
		res = append(res, domain.Core{ID: val.ID, IdUser: val.IdUser, IdContent: val.IdContent, Comment: val.Comment})
	}

	return res
}
