package di

import (
	"github.com/yoshiyoshiharu/item-throw-ways/domain/repository"
	service "github.com/yoshiyoshiharu/item-throw-ways/domain/service/batch"
	handler "github.com/yoshiyoshiharu/item-throw-ways/handler/batch"
	"gorm.io/gorm"
)

func InitAreaCollectWeekday(db *gorm.DB) handler.AreaCollectWeekdayBatchHandler {
  ar := repository.NewAreaCollectWeekdayRepository(db)
  kr := repository.NewKindRepository(db)
	s := service.NewAreaCollectWeekdayBatchService(ar, kr)
	h := handler.NewAreaCollectWeekdayBatchHandler(s)

  return h
}