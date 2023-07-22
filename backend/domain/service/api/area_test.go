package service

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/yoshiyoshiharu/item-throw-ways/infrastructure/entity"
	mock_repository "github.com/yoshiyoshiharu/item-throw-ways/mock/domain/repository"
)

func TestAreaService_FindAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockAreaRepository(ctrl)
	mockRepo.EXPECT().FindAll().Return([]*entity.Area{
		{ID: 1, Name: "Area 1"},
		{ID: 2, Name: "Area 2"},
		{ID: 3, Name: "Area 3"},
	})

	areaService := NewAreaService(mockRepo)

	t.Run("FindAllは全てのAreaを返す", func(t *testing.T) {
		areas := areaService.FindAll()

		assert.Equal(t, 3, len(areas))
	})
}

func TestAreaService_FindById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockAreaRepository(ctrl)
	mockRepo.EXPECT().FindById(1).Return(&entity.Area{ID: 1, Name: "Area 1"}, nil)
	mockRepo.EXPECT().FindById(999).Return(nil, errors.New("Area not found"))

	areaService := NewAreaService(mockRepo)

	t.Run("FindByIdは指定したIDのAreaを返す", func(t *testing.T) {
		area, err := areaService.FindById(1)
		assert.NoError(t, err)
		assert.NotNil(t, area)
		assert.Equal(t, 1, area.ID)
	})

	t.Run("FindByIdは指定したIDのAreaが存在しない場合はエラーを返す", func(t *testing.T) {
		nonExistingID := 999
		area, err := areaService.FindById(nonExistingID)
		assert.Error(t, err)
		assert.Nil(t, area)
	})
}
