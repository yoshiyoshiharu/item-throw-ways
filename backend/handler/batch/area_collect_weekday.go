package handler

import service "github.com/yoshiyoshiharu/item-throw-ways/domain/service/batch"

type AreaCollectWeekdayBatchHandler interface {
	UpdateAll()
}

type areaCollectWeekdayBatchHandler struct {
	s service.AreaCollectWeekdayBatchService
}

func NewAreaCollectWeekdayBatchHandler(service service.AreaCollectWeekdayBatchService) AreaCollectWeekdayBatchHandler {
	return &areaCollectWeekdayBatchHandler{
		s: service,
	}
}

func (h *areaCollectWeekdayBatchHandler) UpdateAll() {
	err := h.s.UpdateAll()

	notifySlack(err)
}
