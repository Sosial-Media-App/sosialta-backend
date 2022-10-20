package delivery

import (
	"github.com/Sosial-Media-App/sosialta/features/contents/domain"
)

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
	ID           uint   `json:"id"`
	IdUser       uint   `json:"id_user"`
	Username     string `json:"username"`
	StoryType    string `json:"story_type"`
	StoryDetail  string `json:"story_detail"`
	StoryPicture string `json:"story_picture"`
}

type GetContentResponse struct {
	ID           uint   `json:"id"`
	IdUser       uint   `json:"id_user"`
	Username     string `json:"username"`
	StoryType    string `json:"story_type"`
	StoryDetail  string `json:"story_detail"`
	StoryPicture string `json:"story_picture"`
	MyContent    bool   `json:"my_content"`
	domain.DetailCore
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "update":
		cnv := core.(domain.Core)
		res = UpdateResponse{
			ID: cnv.ID, IdUser: cnv.IdUser, Username: cnv.Username, StoryType: cnv.StoryType,
			StoryDetail: cnv.StoryDetail, StoryPicture: cnv.StoryPicture,
		}
	case "register":
		cnv := core.(domain.Core)
		res = UpdateResponse{
			ID: cnv.ID, IdUser: cnv.IdUser, Username: cnv.Username, StoryType: cnv.StoryType,
			StoryDetail: cnv.StoryDetail, StoryPicture: cnv.StoryPicture,
		}
	case "getdetail":
		cnv := core.(domain.Core)
		res = GetContentResponse{ID: cnv.ID, IdUser: cnv.IdUser, Username: cnv.Username, StoryType: cnv.StoryType,
			StoryDetail: cnv.StoryDetail, StoryPicture: cnv.StoryPicture, DetailCore: cnv.DetailCore}
	}
	return res
}

func ToResponseContent(core interface{}, code string) interface{} {
	var res interface{}
	var arr []GetContentResponse
	val := core.([]domain.Core)
	for _, cnv := range val {
		arr = append(arr, GetContentResponse{ID: cnv.ID, IdUser: cnv.IdUser, Username: cnv.Username, StoryType: cnv.StoryType,
			StoryDetail: cnv.StoryDetail, StoryPicture: cnv.StoryPicture, DetailCore: cnv.DetailCore})
	}
	res = arr
	return res
}
