package di

import (
	"github.com/yoshiyoshiharu/item-throw-ways/domain/repository"
	service "github.com/yoshiyoshiharu/item-throw-ways/domain/service/api"
	handler "github.com/yoshiyoshiharu/item-throw-ways/handler/api"
	"gorm.io/gorm"
)

func InitItem(db *gorm.DB) handler.ItemHandler {
  r := repository.NewItemRepository(db)
	s := service.NewItemService(r)
	h := handler.NewItemHandler(s)

  return h
}
