package services

import (
	"errors"
	"testing"

	"github.com/Sosial-Media-App/sosialta/config"
	"github.com/Sosial-Media-App/sosialta/features/users/domain"
	"github.com/Sosial-Media-App/sosialta/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetUser(t *testing.T) {
	repo := mocks.NewRepository(t)
	returnData := domain.Core{
		ID: 1, Fullname: "Elon Musk", Username: "elonmusk", Email: "elon@tesla.com",
		Password: "qwerty12345", Phone: "08123456789", Dob: "1988-08-19", UserPicture: "www.cryptotimes.io/wp-content/uploads/2022/03/8817.jpg"}

	t.Run("Sukses Get User", func(t *testing.T) {
		repo.On("Get", mock.Anything).Return(returnData, nil).Once()
		srv := New(repo)
		res, err := srv.GetUser(returnData)
		assert.NotNil(t, res)
		assert.Nil(t, err)
		assert.Equal(t, returnData, res)
		repo.AssertExpectations(t)
	})

	t.Run("Database Error", func(t *testing.T) {
		repo.On("Get", mock.Anything).Return(domain.Core{}, errors.New(config.DATABASE_ERROR)).Once()
		srv := New(repo)
		res, err := srv.GetUser(domain.Core{})
		assert.NotNil(t, res)
		assert.Nil(t, err)
		assert.EqualError(t, errors.New("error, database could not process"), config.DATABASE_ERROR)
		repo.AssertExpectations(t)
	})

	t.Run("No Data", func(t *testing.T) {
		repo.On("Get", mock.Anything).Return(domain.Core{}, errors.New(config.NO_DATA)).Once()
		srv := New(repo)
		res, err := srv.GetUser(domain.Core{})
		assert.NotNil(t, res)
		assert.Nil(t, err)
		assert.EqualError(t, errors.New("no data in database"), config.NO_DATA)
		repo.AssertExpectations(t)
	})
}

func TestAddUser(t *testing.T) {
	repo := mocks.NewRepository(t)
	returnData := domain.Core{
		ID: 1, Fullname: "Elon Musk", Username: "elonmusk", Email: "elon@tesla.com", Password: "qwerty12345"}

	t.Run("Sukses Insert User", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(returnData, nil).Once()
		srv := New(repo)
		input := domain.Core{
			Fullname: "Elon Musk", Username: "elonmusk", Email: "elon@tesla.com", Password: "qwerty12345"}
		res, err := srv.AddUser(input)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, returnData, res)
		repo.AssertExpectations(t)
	})

	t.Run("Encrypt Error", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(domain.Core{}, errors.New(config.ENCRYPT_ERROR)).Once()
		srv := New(repo)
		input := domain.Core{}
		res, err := srv.AddUser(input)
		assert.NotNil(t, res)
		assert.NotNil(t, err)
		assert.EqualError(t, err, config.ENCRYPT_ERROR)
		assert.Equal(t, domain.Core{}, res)
		repo.AssertExpectations(t)
	})

	t.Run("Duplicated Data", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(domain.Core{}, errors.New(config.DUPLICATED_DATA)).Once()
		srv := New(repo)
		input := domain.Core{}
		res, err := srv.AddUser(input)
		assert.NotNil(t, res)
		assert.NotNil(t, err)
		assert.EqualError(t, errors.New("error, data duplicate in database"), config.DUPLICATED_DATA)
		assert.Equal(t, domain.Core{}, res)
		repo.AssertExpectations(t)
	})
}

func TestLogin(t *testing.T) {
	repo := mocks.NewRepository(t)
	returnData := domain.Core{ID: 1, Fullname: "Elon Musk", Username: "elonmusk", Email: "elon@tesla.com", Password: "$2a$10$40zy5ax5eMNDo/SSVmTem.5eqT0bu350nkE6qxEhpOLGgXQ4DXtcS"}

	t.Run("Sukses Login", func(t *testing.T) {
		repo.On("Login", mock.Anything).Return(returnData, nil).Once()
		srv := New(repo)
		input := domain.Core{Email: "elon@tesla.com", Password: "qwerty12345"}
		res, err := srv.Login(input)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, returnData, res)
		repo.AssertExpectations(t)
	})

	t.Run("Database Error", func(t *testing.T) {
		repo.On("Login", mock.Anything).Return(domain.Core{}, errors.New(config.DATABASE_ERROR)).Once()
		srv := New(repo)
		res, err := srv.Login(domain.Core{})
		assert.NotNil(t, res)
		assert.NotNil(t, err)
		assert.EqualError(t, errors.New("error, database could not process"), config.DATABASE_ERROR)
		repo.AssertExpectations(t)
	})

	t.Run("No Data", func(t *testing.T) {
		repo.On("Login", mock.Anything).Return(domain.Core{}, errors.New(config.NO_DATA)).Once()
		srv := New(repo)
		res, err := srv.Login(domain.Core{})
		assert.NotNil(t, res)
		assert.NotNil(t, err)
		assert.EqualError(t, errors.New("no data in database"), config.NO_DATA)
		repo.AssertExpectations(t)
	})
}

func TestUpdateUser(t *testing.T) {
	repo := mocks.NewRepository(t)
	returnData := domain.Core{
		ID: 1, Fullname: "Elon Musk", Username: "elonmusk", Email: "elon@tesla.com", Password: "qwerty12345",
		Phone: "08123456789", Dob: "1988-08-19", UserPicture: "www.cryptotimes.io/wp-content/uploads/2022/03/8817.jpg"}

	t.Run("Sukses Update User", func(t *testing.T) {
		repo.On("Insert", mock.Anything, mock.Anything).Return(returnData, nil).Once()
		srv := New(repo)
		input := domain.Core{
			Fullname: "Elon Musk", Username: "elonmusk", Email: "elon@tesla.com", Password: "qwerty12345",
			Phone: "08123456789", Dob: "1988-08-19", UserPicture: "www.cryptotimes.io/wp-content/uploads/2022/03/8817.jpg"}
		res, err := srv.UpdateUser(input, input.ID)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, returnData, res)
		repo.AssertExpectations(t)
	})
}

func TestDeleteUser(t *testing.T) {
	repo := mocks.NewRepository(t)
	input := domain.Core{}

	t.Run("Sukses Delete User", func(t *testing.T) {
		repo.On("Delete User", mock.Anything).Return(nil).Once()
		srv := New(repo)
		id := input.ID
		err := srv.DeleteUser(id)
		assert.Nil(t, err)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
}
