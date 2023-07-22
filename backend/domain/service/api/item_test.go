package service

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/yoshiyoshiharu/item-throw-ways/infrastructure/entity"
	mock_repository "github.com/yoshiyoshiharu/item-throw-ways/mock/domain/repository"
)

func TestItemService_FindAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockItemRepository(ctrl)
	itemService := NewItemService(mockRepo)

	t.Run("全件返す", func(t *testing.T) {
		mockRepo.EXPECT().FindAll().Return([]*entity.Item{
			{ID: 1, Name: "Item 1"},
			{ID: 2, Name: "Item 2"},
			{ID: 3, Name: "Item 3"},
		})

		items := itemService.FindAll()

		assert.Equal(t, 3, len(items))
	})
}

func TestItemService_DeleteAndInsertAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockItemRepository(ctrl)
	itemService := NewItemService(mockRepo)

	t.Run("成功したときエラーを返さない", func(t *testing.T) {
		mockItems := []*entity.Item{
			{ID: 1, Name: "Item 1"},
			{ID: 2, Name: "Item 2"},
			{ID: 3, Name: "Item 3"},
		}

		mockRepo.EXPECT().DeleteAndInsertAll(mockItems).Return(nil)

		err := itemService.DeleteAndInsertAll(mockItems)

		assert.NoError(t, err)
	})

	t.Run("失敗したときエラーを返す", func(t *testing.T) {
		mockItems := []*entity.Item{
			{ID: 1, Name: "Item 1"},
			{ID: 2, Name: "Item 2"},
			{ID: 3, Name: "Item 3"},
		}

		expectedError := errors.New("delete and insert error")
		mockRepo.EXPECT().DeleteAndInsertAll(mockItems).Return(expectedError)

		err := itemService.DeleteAndInsertAll(mockItems)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}
