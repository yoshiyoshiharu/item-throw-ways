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

func TestItemHandler_FindAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockItemService := mock_service.NewMockItemService(ctrl)
	handler := NewItemHandler(mockItemService)

	kinds := []*entity.Kind{
		entity.NewKind(1, "可燃ごみ"),
		entity.NewKind(2, "不燃ごみ"),
		entity.NewKind(3, "資源"),
		entity.NewKind(4, "粗大ごみ"),
	}

	items := []*entity.Item{
		entity.NewItem(1, "アイロン", "あいろん", 100, "備考1", []entity.Kind{*kinds[0], *kinds[1]}),
		entity.NewItem(1, "鍵", "かぎ", 0, "備考2", []entity.Kind{*kinds[1]}),
	}

	mockItemService.EXPECT().FindAll().Return(items)

  w := httptest.NewRecorder()
  c, _ := gin.CreateTestContext(w)
  c.Request = httptest.NewRequest("GET", "/areas", nil)

	handler.FindAll(c)

  resp, _ := strconv.Unquote(w.Body.String())

	t.Run("[正常系] エリア一覧をJSONで返すこと", func(t *testing.T) {
		assert.Equal(t, 200, w.Code)
		assert.Equal(
      t,
      `[{"id":1,"name":"アイロン","name_kana":"あいろん","price":100,"remarks":"備考1","kinds":[{"id":1,"name":"可燃ごみ"},{"id":2,"name":"不燃ごみ"}]},{"id":1,"name":"鍵","name_kana":"かぎ","price":0,"remarks":"備考2","kinds":[{"id":2,"name":"不燃ごみ"}]}]`,
      resp,
    )
	})
}
