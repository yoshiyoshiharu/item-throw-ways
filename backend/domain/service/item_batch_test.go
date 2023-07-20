package service

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/yoshiyoshiharu/item-throw-ways/infrastructure/entity"
	mock_repository "github.com/yoshiyoshiharu/item-throw-ways/mock/domain/repository"
)

func MockCSVServer() *httptest.Server {
  csvData := `50音,品名,種別,料金,説明
              あ,アイロン,不燃ごみ,,
              あ,アイロン台,粗大ごみ,400,
              あ,アイロンプリント紙,可燃ごみ,,
              あ,空き箱（紙製）,資源,,つぶして、雑誌と一緒にしばって出すか、ひもでしばって出してください。`
  handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/csv")
		io.WriteString(w, csvData)
	}

	return httptest.NewServer(http.HandlerFunc(handler))
}

func TestUpdateAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

  mockKindRepository := mock_repository.NewMockKindRepository(ctrl)
  mockItemRepository := mock_repository.NewMockItemRepository(ctrl)
  itemBatchService := NewItemBatchService(mockItemRepository, mockKindRepository)

	mockAPIServer := MockCSVServer()
	defer mockAPIServer.Close()

  kinds := []*entity.Kind{
		{ID: 1, Name: "可燃ごみ"},
		{ID: 2, Name: "不燃ごみ"},
		{ID: 3, Name: "資源"},
  }

	mockKindRepository.EXPECT().FindAll().Return(kinds)

  expectInsertedItems := []*entity.Item {
    {ID: 1, Name: "アイロン", Kinds: []entity.Kind{*kinds[1]}},
    {ID: 2, Name: "アイロン台", Kinds: []entity.Kind{*kinds[0]}},
    {ID: 3, Name: "アイロンプリント紙", Kinds: []entity.Kind{*kinds[2]}},
    {ID: 4, Name: "空き箱（紙製）", Kinds: []entity.Kind{*kinds[2]}},
  }

	mockItemRepository.EXPECT().DeleteAndInsertAll(expectInsertedItems).Return(nil)

	err := itemBatchService.UpdateAll(mockAPIServer.URL)
  assert.NoError(t, err)
}

