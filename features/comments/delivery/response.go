package delivery

import "github.com/Sosial-Media-App/sosialta/features/comments/domain"

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
	ID        uint   `json:"id"`
	IdUser    uint   `json:"id_user"`
	Username  string `json:"username"`
	IdContent uint   `json:"id_content"`
	Comment   string `json:"comment"`
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "update":
		cnv := core.(domain.Core)
		res = UpdateResponse{
			ID: cnv.ID, IdUser: cnv.IdUser, Username: cnv.Username, IdContent: cnv.IdContent,
			Comment: cnv.Comment,
		}
	case "register":
		cnv := core.(domain.Core)
		res = UpdateResponse{
			ID: cnv.ID, IdUser: cnv.IdUser, Username: cnv.Username, IdContent: cnv.IdContent,
			Comment: cnv.Comment,
		}
	}
	return res
}

func ToResponseComment(core interface{}, code string) interface{} {
	var res interface{}
	var arr []UpdateResponse
	val := core.([]domain.Core)
	for _, cnv := range val {
		arr = append(arr, UpdateResponse{ID: cnv.ID, IdUser: cnv.IdUser, Username: cnv.Username, IdContent: cnv.IdContent,
			Comment: cnv.Comment})
	}
	res = arr
	return res
}
