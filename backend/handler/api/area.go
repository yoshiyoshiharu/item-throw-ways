package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	service "github.com/yoshiyoshiharu/item-throw-ways/domain/service/api"
)

type AreaHandler interface {
	FindAll(*gin.Context)
}

type areaHandler struct {
	s service.AreaService
}

func NewAreaHandler(service service.AreaService) AreaHandler {
	return &areaHandler{
		s: service,
	}
}

func (h *areaHandler) FindAll(c *gin.Context) {
	areas := h.s.FindAll()

	jsonBody, err := json.Marshal(areas)
	if err != nil {
    c.IndentedJSON(http.StatusInternalServerError, err)
	}

  c.IndentedJSON(http.StatusOK, string(jsonBody))
}
