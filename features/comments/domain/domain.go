package domain

import "github.com/labstack/echo/v4"

type Core struct {
	ID        uint
	IdUser    uint
	IdContent uint
	Comment   string
}

type Repository interface {
	Insert(newComment Core) (Core, error)
	Update(newComment Core) (Core, error)
	Delete(id uint) error
	Get(id_content uint) ([]Core, error)
}

type Services interface {
	GetComment(id_content uint) ([]Core, error)
	AddComment(newComment Core) (Core, error)
	UpdateComment(newComment Core) (Core, error)
	DeleteComment(id uint) error
	ExtractToken(c echo.Context) uint
}

type Handler interface {
	AddComment() echo.HandlerFunc
	UpdateComment() echo.HandlerFunc
	GetComment() echo.HandlerFunc
	DeactiveComment() echo.HandlerFunc
}
