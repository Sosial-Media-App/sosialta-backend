package delivery

import "github.com/Sosial-Media-App/sosialta/features/contents/domain"

type RegiterFormat struct {
	IdUser       uint   `json:"id_user" form:"id_user"`
	StoryType    string `json:"story_type" form:"story_type"`
	StoryDetail  string `json:"story_detail" form:"story_detail"`
	StoryPicture string `json:"story_picture" form:"story_picture"`
}

type UpdateFormat struct {
	ID           uint   `json:"id" form:"id"`
	IdUser       uint   `json:"id_user" form:"id_user"`
	StoryType    string `json:"story_type" form:"story_type"`
	StoryDetail  string `json:"story_detail" form:"story_detail"`
	StoryPicture string `json:"story_picture" form:"story_picture"`
}

func ToDomain(i interface{}) domain.Core {
	switch i.(type) {
	case UpdateFormat:
		cnv := i.(UpdateFormat)
		return domain.Core{
			ID:           cnv.ID,
			IdUser:       cnv.IdUser,
			StoryType:    cnv.StoryType,
			StoryDetail:  cnv.StoryDetail,
			StoryPicture: cnv.StoryPicture,
		}

	case RegiterFormat:
		cnv := i.(RegiterFormat)
		return domain.Core{
			IdUser:       cnv.IdUser,
			StoryType:    cnv.StoryType,
			StoryDetail:  cnv.StoryDetail,
			StoryPicture: cnv.StoryPicture,
		}
	}
	return domain.Core{}
}
