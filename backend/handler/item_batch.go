package handler

import (
	"github.com/yoshiyoshiharu/item-throw-ways/domain/service"
)

type ItemBatchHandler interface {
  UpdateAll()
}

type itemBatchHandler struct {
  s service.ItemBatchService
}

func NewItemBatchHandler(service service.ItemBatchService) *itemBatchHandler {
  return &itemBatchHandler{
    s: service,
  }
}

const  API_URL = "https://www.city.bunkyo.lg.jp/library/opendata-bunkyo/01tetsuduki-kurashi/06bunbetuhinmoku/bunbetuhinmoku.csv"

func (h *itemBatchHandler) UpdateAll () {
  h.s.UpdateAll(API_URL)
}
