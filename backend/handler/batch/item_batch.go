package handler

import service "github.com/yoshiyoshiharu/item-throw-ways/domain/service/batch"

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
  
  // notifySlack(err)
}

