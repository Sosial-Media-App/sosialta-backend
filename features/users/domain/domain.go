package domain

import "github.com/labstack/echo/v4"

type Core struct {
	ID          uint
	Username    string
	Email       string
	Password    string
	Fullname    string
	Phone       string
	Dob         string
	UserPicture string
}

type Repository interface {
	Insert(newUser Core) (Core, error)
	Login(newUser Core) (Core, error)
	Update(updateData Core) (Core, error)
	Delete(newUser Core) (Core, error)
}

type Services interface {
	AddUser(newUser Core) (Core, error)
	Login(newUser Core) (Core, error)
	UpdateUser(updateData Core) (Core, error)
	DeleteUser(newUser Core) error
	GenerateToken(id uint, username string) string
	ExtractToken(c echo.Context) uint
}

type Handler interface {
	RegiterUser() echo.HandlerFunc
	LoginUser() echo.HandlerFunc
	UpdateDataUser() echo.HandlerFunc
}
