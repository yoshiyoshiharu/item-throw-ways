package di

import (
	"github.com/yoshiyoshiharu/item-throw-ways/domain/repository"
	service "github.com/yoshiyoshiharu/item-throw-ways/domain/service/api"
	handler "github.com/yoshiyoshiharu/item-throw-ways/handler/api"
	"gorm.io/gorm"
)

func InitAreaCollectDate(db *gorm.DB) handler.AreaCollectDateHandler {
  r := repository.NewAreaCollectWeekdayRepository(db)
	s := service.NewAreaCollectWeekdayService(r)
	h := handler.NewAreaCollectDateHandler(s)

  return h
}
