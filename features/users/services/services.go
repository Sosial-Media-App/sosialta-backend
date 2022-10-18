package services

import (
	"errors"
	"strings"

	"github.com/Sosial-Media-App/sosialta/config"
	"github.com/Sosial-Media-App/sosialta/features/users/domain"
)

type userService struct {
	qry domain.Repository
}

func New(repo domain.Repository) domain.Services {
	return &userService{
		qry: repo,
	}
}

func (us *userService) AddUser(newUser domain.Core) (domain.Core, error) {

	return newUser, nil
}

func (us *userService) Login(newUser domain.Core) (domain.Core, error) {

	return newUser, nil
}

func (us *userService) UpdateUser(updateData domain.Core) (domain.Core, error) {
	res, err := us.qry.Update(updateData)
	if err != nil {
		if strings.Contains(err.Error(), config.DUPLICATED_DATA) {
			return domain.Core{}, errors.New(config.REJECTED_DATA)
		}
	}

	return res, nil
}

func (us *userService) DeleteUser(newUser domain.Core) error {

	return nil
}
