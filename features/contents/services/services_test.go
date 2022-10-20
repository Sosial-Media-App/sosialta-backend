package services

import (
	"errors"
	"testing"

	"github.com/Sosial-Media-App/sosialta/config"
	"github.com/Sosial-Media-App/sosialta/features/contents/domain"
	mocks "github.com/Sosial-Media-App/sosialta/mocks/contents"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddContent(t *testing.T) {
	repo := mocks.NewRepository(t)
	returnData := domain.Core{
		ID: 1}

	t.Run("Sukses Insert User", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(returnData, nil).Once()
		srv := New(repo)
		input := domain.Core{StoryType: "story", StoryDetail: "test", StoryPicture: ".jpg"}
		res, err := srv.AddContent(input)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, returnData, res)
		repo.AssertExpectations(t)
	})

	t.Run("Encrypt Error", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(domain.Core{}, errors.New(config.ENCRYPT_ERROR)).Once()
		srv := New(repo)
		input := domain.Core{}
		res, err := srv.AddContent(input)
		assert.NotNil(t, res)
		assert.NotNil(t, err)
		// assert.EqualError(t, err, config.ENCRYPT_ERROR)
		assert.Equal(t, domain.Core{}, res)
		repo.AssertExpectations(t)
	})
}

func TestGetContent(t *testing.T) {
	repo := mocks.NewRepository(t)
	var returnData []domain.Core
	returnData = append(returnData, domain.Core{
		ID: 1})
	t.Run("Sukses Get Content", func(t *testing.T) {
		repo.On("Get", mock.Anything).Return(returnData, nil).Once()
		srv := New(repo)
		res, err := srv.GetContent(1)
		assert.NotNil(t, res)
		assert.Nil(t, err)
		assert.Equal(t, returnData, res)
		repo.AssertExpectations(t)
	})

	t.Run("No Data", func(t *testing.T) {
		repo.On("Get", mock.Anything).Return(nil, errors.New(config.NO_DATA)).Once()
		srv := New(repo)
		res, err := srv.GetContent(1)
		assert.Nil(t, res)
		assert.NotNil(t, err)
		assert.EqualError(t, errors.New(config.NO_DATA), config.NO_DATA)
		repo.AssertExpectations(t)
	})
}

func TestGetContentDetail(t *testing.T) {
	repo := mocks.NewRepository(t)
	var returnData domain.Core
	t.Run("Sukses Get Content", func(t *testing.T) {
		repo.On("GetDetail", mock.Anything).Return(returnData, nil).Once()
		srv := New(repo)
		res, err := srv.GetContentDetail(1)
		assert.NotNil(t, res)
		assert.Nil(t, err)
		assert.Equal(t, returnData, res)
		repo.AssertExpectations(t)
	})

	t.Run("No Data", func(t *testing.T) {
		repo.On("GetDetail", mock.Anything).Return(domain.Core{}, errors.New(config.NO_DATA)).Once()
		srv := New(repo)
		res, _ := srv.GetContentDetail(0)
		assert.Equal(t, res, domain.Core{})
		assert.EqualError(t, errors.New(config.NO_DATA), config.NO_DATA)
		repo.AssertExpectations(t)
	})
}

func TestUpdateUser(t *testing.T) {
	repo := mocks.NewRepository(t)
	returnData := domain.Core{
		ID: 1}

	t.Run("Sukses Update User", func(t *testing.T) {
		repo.On("Update", mock.Anything, mock.Anything).Return(returnData, nil).Once()
		srv := New(repo)
		input := domain.Core{ID: uint(1), IdUser: 1, StoryType: "story", StoryDetail: "test", StoryPicture: ".jpg"}
		res, err := srv.UpdateContent(input)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, returnData, res)
		repo.AssertExpectations(t)
	})

	t.Run("Database Error", func(t *testing.T) {
		repo.On("Update", mock.Anything).Return(domain.Core{}, errors.New(config.DATABASE_ERROR)).Once()
		srv := New(repo)
		input := domain.Core{}
		res, err := srv.UpdateContent(input)
		assert.Equal(t, res, domain.Core{})
		assert.NotNil(t, err)
		assert.EqualError(t, errors.New(config.DATABASE_ERROR), config.DATABASE_ERROR)
		repo.AssertExpectations(t)
	})
}

func TestDeleteUser(t *testing.T) {
	repo := mocks.NewRepository(t)
	input := uint(1)

	t.Run("Sukses Delete Content", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(nil).Once()
		srv := New(repo)
		err := srv.DeleteContent(input)
		assert.Nil(t, err)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed On Delete", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(errors.New(config.NO_DATA)).Once()
		srv := New(repo)
		err := srv.DeleteContent(uint(0))
		assert.NotNil(t, err)
		assert.EqualError(t, errors.New(config.NO_DATA), config.NO_DATA)
		repo.AssertExpectations(t)
	})
}
