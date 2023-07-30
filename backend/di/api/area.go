package di

import (
	"github.com/yoshiyoshiharu/item-throw-ways/domain/repository"
	service "github.com/yoshiyoshiharu/item-throw-ways/domain/service/api"
	handler "github.com/yoshiyoshiharu/item-throw-ways/handler/api"
	"gorm.io/gorm"
)

func InitArea(db *gorm.DB) handler.AreaHandler {
  r := repository.NewAreaRepository(db)
	s := service.NewAreaService(r)
	h := handler.NewAreaHandler(s)

  return h
}
