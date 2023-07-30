package handler

import (
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/yoshiyoshiharu/item-throw-ways/infrastructure/entity"
	mock_service "github.com/yoshiyoshiharu/item-throw-ways/mock/domain/service/api"
)

func TestAreaHandler_FindAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAreaService := mock_service.NewMockAreaService(ctrl)
	handler := NewAreaHandler(mockAreaService)

	areas := []*entity.Area{
		entity.NewArea(1, "Area 1"),
		entity.NewArea(2, "Area 2"),
	}

	mockAreaService.EXPECT().FindAll().Return(areas)

  w := httptest.NewRecorder()
  c, _ := gin.CreateTestContext(w)
  c.Request = httptest.NewRequest("GET", "/areas", nil)

	handler.FindAll(c)

  resp, _ := strconv.Unquote(w.Body.String())

	t.Run("[正常系] エリア一覧をJSONで返すこと", func(t *testing.T) {
		assert.Equal(t, 200, w.Code)
		assert.Equal(
      t,
      `[{"id":1,"name":"Area 1"},{"id":2,"name":"Area 2"}]`,
      resp,
    )
	})
}
