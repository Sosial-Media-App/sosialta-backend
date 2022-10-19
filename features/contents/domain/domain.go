package domain

import "github.com/labstack/echo/v4"

type Core struct {
	ID           uint
	IdUser       uint
	StoryType    string
	StoryDetail  string
	StoryPicture string
	DetailCore
}

type CommentCore struct {
	ID        uint
	IdUser    uint
	IdContent uint
	Comment   string
}

type DetailCore struct {
	DetailCore []CommentCore
}

type Repository interface {
	Insert(newContent Core) (Core, error)
	Update(updateData Core) (Core, error)
	Delete(id uint) error
	Get() ([]Core, error)
	GetDetail(newContent Core) (Core, error)
}

type Services interface {
	GetContent() ([]Core, error)
	GetContentDetail(newContent Core) (Core, error)
	AddContent(newContent Core) (Core, error)
	UpdateContent(updateData Core) (Core, error)
	DeleteContent(id uint) error
	ExtractToken(c echo.Context) uint
}

type Handler interface {
	RegiterContent() echo.HandlerFunc
	UpdateDataContent() echo.HandlerFunc
	GetContent() echo.HandlerFunc
	DeactiveContent() echo.HandlerFunc
}
