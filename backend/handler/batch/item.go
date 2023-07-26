package handler

import (
	"os"

	service "github.com/yoshiyoshiharu/item-throw-ways/domain/service/batch"
)

type ItemBatchHandler interface {
	UpdateAll()
}

type itemBatchHandler struct {
	s service.ItemBatchService
}

func NewItemBatchHandler(service service.ItemBatchService) ItemBatchHandler {
	return &itemBatchHandler{
		s: service,
	}
}

func (h *itemBatchHandler) UpdateAll() {
	err := h.s.UpdateAll()

  if os.Getenv("ENV") == "production" {
    notifySlack(err)
  }
}
