package di

import (
	"github.com/yoshiyoshiharu/item-throw-ways/domain/repository"
	service "github.com/yoshiyoshiharu/item-throw-ways/domain/service/batch"
	handler "github.com/yoshiyoshiharu/item-throw-ways/handler/batch"
	"gorm.io/gorm"
)

func InitItemBatch(db *gorm.DB) handler.ItemBatchHandler {
	kr := repository.NewKindRepository(db)
	ir := repository.NewItemRepository(db)
	s := service.NewItemBatchService(ir, kr)
  h := handler.NewItemBatchHandler(s)

  return h
}
