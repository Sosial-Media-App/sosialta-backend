package delivery

import "github.com/Sosial-Media-App/sosialta/features/comments/domain"

type RegiterFormat struct {
	IdUser    uint   `json:"id_user" form:"id_user"`
	Username  string `json:"username" form:"username"`
	IdContent uint   `json:"id_content" form:"id_content"`
	Comment   string `json:"comment" form:"comment"`
}
type UpdateFormat struct {
	ID        uint   `json:"id" form:"id"`
	IdUser    uint   `json:"id_user" form:"id_user"`
	Username  string `json:"username" form:"username"`
	IdContent uint   `json:"id_content" form:"id_content"`
	Comment   string `json:"comment" form:"comment"`
}

func ToDomain(i interface{}) domain.Core {
	switch i.(type) {
	case UpdateFormat:
		cnv := i.(UpdateFormat)
		return domain.Core{
			ID:        cnv.ID,
			IdUser:    cnv.IdUser,
			Username:  cnv.Username,
			IdContent: cnv.IdContent,
			Comment:   cnv.Comment,
		}
	case RegiterFormat:
		cnv := i.(RegiterFormat)
		return domain.Core{
			IdUser:    cnv.IdUser,
			Username:  cnv.Username,
			IdContent: cnv.IdContent,
			Comment:   cnv.Comment,
		}
	}
	return domain.Core{}
}
