package services

import "github.com/Sosial-Media-App/sosialta/features/users/domain"

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

func (us *userService) UpdateUser(newUser domain.Core) (domain.Core, error) {

	return newUser, nil
}

func (us *userService) DeleteUser(newUser domain.Core) error {

	return nil
}
