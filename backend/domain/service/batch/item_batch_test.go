package service

import (
	"encoding/csv"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/yoshiyoshiharu/item-throw-ways/infrastructure/entity"
	mock_repository "github.com/yoshiyoshiharu/item-throw-ways/mock/domain/repository"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func mockItemCsvServer() *httptest.Server {
	csvData := [][]string{
		{"50音", "品名", "種別", "料金", "説明"},
		{"あ", "アイロン", "不燃ごみ", "100円", ""},
		{"い", "板類一束", "可燃ごみ、粗大ごみ", "400円", "釘は必ず抜いてください。おおむね一辺が30㎝以上のものは粗大ごみへ。"},
	}

	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/csv")
		w.WriteHeader(http.StatusOK)
		csvWriter := csv.NewWriter(transform.NewWriter(w, japanese.ShiftJIS.NewEncoder()))
		csvWriter.WriteAll(csvData)
		csvWriter.Flush()
	}))

	return s
}

func mockHiraganaTranslationServer() *httptest.Server {
	resp := ResponseBody{
		Converted: "ひらがな",
	}

	jsonData, _ := json.Marshal(resp)

	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
	}))

	return s

}
func TestItemBatch_UpdateAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockItemCsvServer := mockItemCsvServer()
	defer mockItemCsvServer.Close()

	mockHiraganaTranslationServer := mockHiraganaTranslationServer()

	ItemsApiUrl = mockItemCsvServer.URL
	HiraganaTranslationApiUrl = mockHiraganaTranslationServer.URL

	kr := mock_repository.NewMockKindRepository(ctrl)
	ir := mock_repository.NewMockItemRepository(ctrl)
	service := NewItemBatchService(ir, kr)

	allKinds := []*entity.Kind{
		entity.NewKind(1, "可燃ごみ"),
		entity.NewKind(2, "不燃ごみ"),
		entity.NewKind(3, "資源"),
		entity.NewKind(4, "粗大ごみ"),
	}
	items := []entity.Item{
		*entity.NewItem(1, "アイロン", "ひらがな", 100, "", []entity.Kind{*allKinds[1]}),
		*entity.NewItem(1, "板類人束", "ひらがな", 400, "釘は必ず抜いてください。おおむね一辺が30㎝以上のものは粗大ごみへ。", []entity.Kind{*allKinds[0], *allKinds[3]}),
	}

	kr.EXPECT().FindAll().Return(allKinds).Times(1)
	ir.EXPECT().DeleteAndInsertAll(items).Return(nil).Times(1)

	err := service.UpdateAll()

	assert.NoError(t, err)
}
