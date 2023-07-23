package service

import (
	"encoding/csv"
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

func mockAreaCollectWeekdayCsvServer() *httptest.Server {
	csvData := [][]string{
		{"町名", "丁目", "可燃ごみ", "不燃ごみ", "資源ごみ"},
		{"後楽", "1丁目", "火曜日・金曜日", "第1・第3木曜日", "水曜日"},
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

func TestAreaCollectWeekdayBatch_UpdateAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockServer := mockAreaCollectWeekdayCsvServer()
	defer mockServer.Close()

	AreaCollectWeekdaysApiUrl = mockServer.URL

	kr := mock_repository.NewMockKindRepository(ctrl)
	ar := mock_repository.NewMockAreaCollectWeekdayRepository(ctrl)
	service := NewAreaCollectWeekdayBatchService(ar, kr)

	allKinds := []*entity.Kind{
		entity.NewKind(1, "可燃ごみ"),
		entity.NewKind(2, "不燃ごみ"),
		entity.NewKind(3, "資源"),
	}
	area := entity.NewArea(1, "後楽1丁目")
	areaCollectWeekdays := []entity.AreaCollectWeekday{
		*entity.NewAreaCollectWeekday(area, allKinds[0], 2, 0),
		*entity.NewAreaCollectWeekday(area, allKinds[0], 5, 0),
		*entity.NewAreaCollectWeekday(area, allKinds[1], 4, 1),
		*entity.NewAreaCollectWeekday(area, allKinds[1], 4, 3),
		*entity.NewAreaCollectWeekday(area, allKinds[2], 3, 0),
	}

	kr.EXPECT().FindAll().Return(allKinds).Times(1)
	ar.EXPECT().DeleteAndInsertAll(areaCollectWeekdays).Return(nil).Times(1)

	err := service.UpdateAll()

	assert.NoError(t, err)
}
