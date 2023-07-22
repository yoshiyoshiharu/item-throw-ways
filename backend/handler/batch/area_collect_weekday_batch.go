package handler

import (
	"github.com/yoshiyoshiharu/item-throw-ways/domain/service"
)
type AreaCollectWeekdayBatchHandler interface {
  UpdateAll()
}

type areaCollectWeekdayBatchHandler struct {
  s service.AreaCollectWeekdayBatchService
}

func NewAreaCollectWeekdayBatchHandler(service service.AreaCollectWeekdayBatchService) *areaCollectWeekdayBatchHandler {
  return &areaCollectWeekdayBatchHandler{
    s: service,
  }
}

func (h *areaCollectWeekdayBatchHandler) UpdateAll () {
  err := h.s.UpdateAll()

  notifySlack(err)
}

