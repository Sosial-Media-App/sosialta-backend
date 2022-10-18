package repository

import (
	"github.com/Sosial-Media-App/sosialta/features/users/domain"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string `gorm:"uniqueIndex:idx_username"`
	Email       string `gorm:"uniqueIndex:idx_email"`
	Password    string
	Fullname    string
	Phone       string
	Dob         string
	UserPicture string
}

func FromDomain(du domain.Core) User {
	return User{
		Model:       gorm.Model{ID: du.ID},
		Username:    du.Username,
		Email:       du.Email,
		Password:    du.Password,
		Fullname:    du.Fullname,
		Phone:       du.Phone,
		Dob:         du.Dob,
		UserPicture: du.UserPicture,
	}
}

func ToDomain(u User) domain.Core {
	return domain.Core{
		ID:          u.ID,
		Username:    u.Username,
		Email:       u.Email,
		Password:    u.Password,
		Fullname:    u.Fullname,
		Phone:       u.Phone,
		Dob:         u.Dob,
		UserPicture: u.UserPicture,
	}
}

func ToDomainArray(u []User) []domain.Core {
	var res []domain.Core
	for _, val := range u {
		res = append(res, domain.Core{ID: val.ID, Username: val.Username, Email: val.Email,
			Password: val.Password, Fullname: val.Fullname, Phone: val.Phone, Dob: val.Dob, UserPicture: val.UserPicture})
	}
	return res
}
