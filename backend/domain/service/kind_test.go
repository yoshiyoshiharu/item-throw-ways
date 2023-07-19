package service

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/yoshiyoshiharu/item-throw-ways/infrastructure/entity"
	mock_repository "github.com/yoshiyoshiharu/item-throw-ways/mock/domain/repository"
)

func TestKindService_FindAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockKindRepository(ctrl)
	mockRepo.EXPECT().FindAll().Return([]*entity.Kind{
		{ID: 1, Name: "Kind 1"},
		{ID: 2, Name: "Kind 2"},
		{ID: 3, Name: "Kind 3"},
	})

	kindService := NewKindService(mockRepo)

  t.Run("FindAllは全てのKindを返す", func(t *testing.T) {
    kinds := kindService.FindAll()

    assert.Equal(t, 3, len(kinds))
  })
}

