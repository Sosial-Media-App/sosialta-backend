package services

import (
	"errors"
	"strings"
	"time"

	"github.com/Sosial-Media-App/sosialta/config"
	"github.com/Sosial-Media-App/sosialta/features/users/domain"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	qry domain.Repository
}

func New(repo domain.Repository) domain.Services {
	return &userService{
		qry: repo,
	}
}

func (us *userService) GetUser(newUser domain.Core) (domain.Core, error) {
	res, err := us.qry.Get(newUser)
	if err != nil {
		log.Error(err.Error())
		if strings.Contains(err.Error(), "table") {
			return domain.Core{}, errors.New("database error")
		} else if strings.Contains(err.Error(), "found") {
			return domain.Core{}, errors.New("no data")
		}
	}

	return res, nil
}

func (us *userService) AddUser(newUser domain.Core) (domain.Core, error) {
	if newUser.Password != "" {
		generate, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)
		newUser.Password = string(generate)
	}
	res, err := us.qry.Insert(newUser)

	if err != nil {
		if newUser.Password == "" {
			return domain.Core{}, errors.New(config.ENCRYPT_ERROR)
		}
		return domain.Core{}, errors.New(config.DUPLICATED_DATA)
	}

	return res, nil
}

func (us *userService) Login(newUser domain.Core) (domain.Core, error) {
	res, err := us.qry.Login(newUser)
	if err != nil {
		log.Error(err.Error())
		if strings.Contains(err.Error(), "table") {
			return domain.Core{}, errors.New("database error")
		} else if strings.Contains(err.Error(), "found") {
			return domain.Core{}, errors.New("no data")
		}
	}
	// token := GenerateToken(res.ID)
	err = bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(newUser.Password))
	if err != nil {
		return domain.Core{}, errors.New("password tidak cocok")
	}
	return res, nil
}

func (us *userService) UpdateUser(updateData domain.Core, id uint) (domain.Core, error) {
	if updateData.Password != "" {
		hashed, _ := bcrypt.GenerateFromPassword([]byte(updateData.Password), bcrypt.DefaultCost)

		updateData.Password = string(hashed)
	}

	res, err := us.qry.Update(updateData, id)
	if err != nil {
		if strings.Contains(err.Error(), config.DUPLICATED_DATA) {
			return domain.Core{}, errors.New(config.REJECTED_DATA)
		}
	}

	return res, nil
}

func (us *userService) DeleteUser(id uint) error {
	err := us.qry.Delete(id)
	if err != nil {
		return errors.New("data not found")
	}
	return nil
}

func (us *userService) GenerateToken(id uint, username string) string {
	claim := make(jwt.MapClaims)
	claim["authorized"] = true
	claim["id"] = id
	claim["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	str, err := token.SignedString([]byte("Sosialta!!!12"))
	if err != nil {
		log.Error(err.Error())
		return ""
	}

	return str
}

func (bs *userService) ExtractToken(c echo.Context) uint {
	token := c.Get("user").(*jwt.Token)
	if token.Valid {
		claim := token.Claims.(jwt.MapClaims)
		return uint(claim["id"].(float64))
	}

	return 0
}
