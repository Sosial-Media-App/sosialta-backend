package repository

import (
	"github.com/Sosial-Media-App/sosialta/features/users/domain"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string `gorm:"unique"`
	Email       string `gorm:"unique"`
	Password    string
	Fullname    string
	Phone       string
	Dob         string
	UserPicture string
	Contents    []Content `gorm:"foreignKey:id"`
}

type Content struct {
	gorm.Model
	IdUser       uint
	StoryType    string
	StroyDetail  string
	StoryPicture string
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

func ToDomainArray(u User, c []Content) domain.Core {
	// var hasil domain.DetailCore
	var cContext []domain.ContentCore

	for _, val := range c {
		cContext = append(cContext, domain.ContentCore{
			ID: val.ID, StoryType: val.StoryType, StroyDetail: val.StroyDetail,
			StoryPicture: val.StoryPicture})
	}
	var res domain.Core = domain.Core{
		ID: u.ID, Username: u.Username, Email: u.Email,
		Password: u.Password, Fullname: u.Fullname, Phone: u.Phone,
		Dob: u.Dob, UserPicture: u.UserPicture, DetailCore: domain.DetailCore{cContext}}

	return res
}
