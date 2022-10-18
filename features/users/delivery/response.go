package delivery

import "github.com/Sosial-Media-App/sosialta/features/users/domain"

func SuccessResponse(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
	}
}

func FailedResponse(msg string) map[string]string {
	return map[string]string{
		"message": msg,
	}
}

type UpdateResponse struct {
	ID          uint   `json:"id"`
	Fullname    string `json:"fullname"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Phone       string `json:"phone"`
	Dob         string `json:"dob"`
	UserPicture string `json:"user_picture"`
	token       string `json:"token"`
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "update":
		cnv := core.(domain.Core)
		res = UpdateResponse{
			ID: cnv.ID, Fullname: cnv.Fullname, Username: cnv.Username,
			Email: cnv.Email, Password: cnv.Password, Phone: cnv.Phone,
			Dob: cnv.Dob, UserPicture: cnv.UserPicture,
		}
	case "register":
		cnv := core.(domain.Core)
		res = UpdateResponse{
			ID: cnv.ID, Fullname: cnv.Fullname, Username: cnv.Username,
			Email: cnv.Email, Password: cnv.Password,
		}
	case "login":
		cnv := core.(domain.Core)
		res = UpdateResponse{
			ID: cnv.ID, Fullname: cnv.Fullname, Username: cnv.Username,
			Email: cnv.Email, Password: cnv.Password,
		}
	}
	return res
}

func ToResponseLogin(core interface{}, userToken string, code string) interface{} {
	var res interface{}
	cnv := core.(domain.Core)
	res = UpdateResponse{
		ID: cnv.ID, Fullname: cnv.Fullname, Username: cnv.Username,
		Email: cnv.Email, Password: cnv.Password, token: userToken,
	}
	return res
}
