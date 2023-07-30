package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	service "github.com/yoshiyoshiharu/item-throw-ways/domain/service/api"
)

type ItemHandler interface {
	FindAll(*gin.Context)
}

type itemHandler struct {
	s service.ItemService
}

func NewItemHandler(service service.ItemService) ItemHandler {
	return &itemHandler{
		s: service,
	}
}

func (h *itemHandler) FindAll(c *gin.Context) {
	items := h.s.FindAll()

	jsonBody, err := json.Marshal(items)
	if err != nil {
    c.IndentedJSON(http.StatusInternalServerError, err)
	}

  c.IndentedJSON(http.StatusOK, string(jsonBody))
}
