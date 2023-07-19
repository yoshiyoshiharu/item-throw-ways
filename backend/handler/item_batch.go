package handler

import (
	"github.com/yoshiyoshiharu/item-throw-ways/domain/service"
)

type ItemBatchHandler interface {
  FindAll()
}

type itemBatchHandler struct {
  s service.ItemBatchService
}

func NewItemBatchHandler(service service.ItemBatchService) *itemBatchHandler {
  return &itemBatchHandler{
    s: service,
  }
}

func (h *itemBatchHandler) UpdateAll () {
  h.s.UpdateAll()
}
