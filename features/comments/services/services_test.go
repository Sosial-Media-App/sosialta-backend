package services

import (
	"errors"
	"testing"

	"github.com/Sosial-Media-App/sosialta/config"
	"github.com/Sosial-Media-App/sosialta/features/comments/domain"
	mocks "github.com/Sosial-Media-App/sosialta/mocks/comments"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetComment(t *testing.T) {
	repo := mocks.NewRepository(t)
	returnData := []domain.Core{{
		ID: 1, IdUser: 1, IdContent: 1, Comment: "Pertamax gan, cendolnya dong!",
	}}

	t.Run("Sukses Get Comment", func(t *testing.T) {
		repo.On("Get", mock.Anything).Return(returnData, nil).Once()
		srv := New(repo)
		input := uint(1)
		res, err := srv.GetComment(input)
		assert.NotNil(t, res)
		assert.Nil(t, err)
		assert.Equal(t, returnData, res)
		repo.AssertExpectations(t)
	})

	t.Run("Database Error", func(t *testing.T) {
		repo.On("Get", mock.Anything).Return(nil, errors.New(config.DATABASE_ERROR)).Once()
		srv := New(repo)
		input := uint(0)
		res, err := srv.GetComment(input)
		assert.Nil(t, res)
		assert.Nil(t, err)
		assert.EqualError(t, errors.New(config.DATABASE_ERROR), config.DATABASE_ERROR)
		repo.AssertExpectations(t)
	})

	t.Run("No Data", func(t *testing.T) {
		repo.On("Get", mock.Anything).Return(nil, errors.New(config.NO_DATA)).Once()
		srv := New(repo)
		input := uint(0)
		res, err := srv.GetComment(input)
		assert.Nil(t, res)
		assert.Nil(t, err)
		assert.EqualError(t, errors.New(config.NO_DATA), config.NO_DATA)
		repo.AssertExpectations(t)
	})
}

func TestAddComment(t *testing.T) {
	repo := mocks.NewRepository(t)
	returnData := domain.Core{
		ID: 1, IdUser: 1, IdContent: 1, Comment: "Pertamax gan, cendolnya dong!"}

	t.Run("Sukses Add Comment", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(returnData, nil).Once()
		srv := New(repo)
		input := domain.Core{IdContent: 1, Comment: "Chronoptic Energy!"}
		res, err := srv.AddComment(input)
		assert.NotNil(t, res)
		assert.Nil(t, err)
		assert.Equal(t, returnData, res)
		repo.AssertExpectations(t)
	})

	t.Run("Duplicated Data", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(domain.Core{}, errors.New(config.DUPLICATED_DATA)).Once()
		srv := New(repo)
		input := domain.Core{IdContent: 1, Comment: "Chronoptic Energy!"}
		res, err := srv.AddComment(input)
		assert.NotNil(t, res)
		assert.NotNil(t, err)
		assert.EqualError(t, errors.New(config.DUPLICATED_DATA), config.DUPLICATED_DATA)
		repo.AssertExpectations(t)
	})
}

func TestUpdateComment(t *testing.T) {
	repo := mocks.NewRepository(t)
	returnData := domain.Core{ID: 1, IdUser: 1, IdContent: 1, Comment: "Pertamax gan, cendolnya dong!"}

	t.Run("Sukses Update Comment", func(t *testing.T) {
		repo.On("Update", mock.Anything, mock.Anything).Return(returnData, nil).Once()
		srv := New(repo)
		// input := domain.Core{ID: 1, Comment: "Chronoptic Energy!"}
		res, err := srv.UpdateComment(returnData)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, returnData, res)
		repo.AssertExpectations(t)
	})

	t.Run("Rejected from Database", func(t *testing.T) {
		repo.On("Update", mock.Anything).Return(domain.Core{}, errors.New(config.REJECTED_DATA)).Once()
		srv := New(repo)
		input := domain.Core{}
		res, _ := srv.UpdateComment(input)
		assert.Equal(t, res, domain.Core{})
		assert.EqualError(t, errors.New(config.REJECTED_DATA), config.REJECTED_DATA)
		repo.AssertExpectations(t)
	})
}

func TestDeleteComment(t *testing.T) {
	repo := mocks.NewRepository(t)
	input := uint(1)
	t.Run("Sukses Delete Comment", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(nil).Once()
		srv := New(repo)
		err := srv.DeleteComment(input)
		assert.Nil(t, err)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Data Not Found", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(errors.New(config.NO_DATA)).Once()
		srv := New(repo)
		err := srv.DeleteComment(uint(0))
		assert.NotNil(t, err)
		assert.EqualError(t, errors.New(config.NO_DATA), config.NO_DATA)
		repo.AssertExpectations(t)
	})
}
