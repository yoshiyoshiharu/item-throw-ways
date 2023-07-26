package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	service "github.com/yoshiyoshiharu/item-throw-ways/domain/service/api"
)

type AreaCollectDateHandler interface {
	FindAll(*gin.Context)
}

type areaCollectDateHandler struct {
	s service.AreaCollectWeekdayService
}

func NewAreaCollectDateHandler(service service.AreaCollectWeekdayService) AreaCollectDateHandler {
	return &areaCollectDateHandler{
		s: service,
	}
}

func (h *areaCollectDateHandler) FindAll(c *gin.Context) {
	area_id, year, month, err := parseParams(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "request parameter is invalid")
	}

	areas := h.s.ConvertByAreaWithAroundMonths(area_id, year, month)

	jsonBody, err := json.Marshal(areas)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	}

  c.IndentedJSON(http.StatusOK, string(jsonBody))
}

func parseParams(c *gin.Context) (int, int, time.Month, error) {
	area_id, err := strconv.Atoi(c.Param("area_id"))
	year, err := strconv.Atoi(c.Param("year"))
	monthInt, err := strconv.Atoi(c.Param("month"))
	month, err := intToMonth(monthInt)

	if err != nil {
		return 0, 0, 0, err
	}
	return area_id, year, month, nil
}

func intToMonth(monthInt int) (time.Month, error) {
	if monthInt < 1 || monthInt > 12 {
		return 0, fmt.Errorf("Invalid month: %d", monthInt)
	}
	return time.Month(monthInt), nil
}
