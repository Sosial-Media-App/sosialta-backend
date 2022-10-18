package delivery

import "github.com/Sosial-Media-App/sosialta/features/users/domain"

type RegiterFormat struct {
	Fullname    string `json:"fullname" form:"fullname"`
	Username    string `json:"username" form:"username"`
	Email       string `json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
	Phone       string `json:"phone" form:"phone"`
	Dob         string `json:"dob" form:"dob"`
	UserPicture string `json:"user_picture" form:"user_picture"`
}
type UpdateFormat struct {
	Fullname    string `json:"fullname" form:"fullname"`
	Username    string `json:"username" form:"username"`
	Email       string `json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
	Phone       string `json:"phone" form:"phone"`
	Dob         string `json:"dob" form:"dob"`
	UserPicture string `json:"user_picture" form:"user_picture"`
}

func ToDomain(i interface{}) domain.Core {
	switch i.(type) {
	case UpdateFormat:
		cnv := i.(UpdateFormat)
		return domain.Core{
			Fullname:    cnv.Fullname,
			Username:    cnv.Username,
			Email:       cnv.Email,
			Password:    cnv.Password,
			Phone:       cnv.Phone,
			Dob:         cnv.Dob,
			UserPicture: cnv.UserPicture,
		}
	case RegiterFormat:
		cnv := i.(RegiterFormat)
		return domain.Core{
			Fullname:    cnv.Fullname,
			Username:    cnv.Username,
			Email:       cnv.Email,
			Password:    cnv.Password,
			Phone:       cnv.Phone,
			Dob:         cnv.Dob,
			UserPicture: cnv.UserPicture,
		}
	}
	return domain.Core{}
}
