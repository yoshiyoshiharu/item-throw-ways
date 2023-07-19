package service

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/yoshiyoshiharu/item-throw-ways/infrastructure/entity"
	mock_repository "github.com/yoshiyoshiharu/item-throw-ways/mock/domain/repository"
)

func TestAreaCollectWeekdayService_GetByAreaWithAroundMonths(t *testing.T) {
  mockCtrl := gomock.NewController(t)
  defer mockCtrl.Finish()

  mockAreaCollectWeekdayRepository := mock_repository.NewMockAreaCollectWeekdayRepository(mockCtrl)

  area := entity.NewArea(1, "後楽一丁目")
  kinds := map[string]*entity.Kind{
    "kanen": entity.NewKind(1, "可燃ごみ"),
    "funen": entity.NewKind(2, "不燃ごみ"),
    "shigen": entity.NewKind(3, "資源"),
  }
  year := 2023
  month := time.Month(7)
  loc := time.FixedZone("Asia/Tokyo", 9*60*60)

  mockAreaCollectWeekdays := []*entity.AreaCollectWeekday{
    entity.NewAreaCollectWeekday(area, kinds["kanen"], 2, 0),
    entity.NewAreaCollectWeekday(area, kinds["kanen"], 5, 0),
    entity.NewAreaCollectWeekday(area, kinds["funen"], 4, 1),
    entity.NewAreaCollectWeekday(area, kinds["funen"], 4, 3),
    entity.NewAreaCollectWeekday(area, kinds["shigen"], 3, 0),
  }

  expected := []*entity.AreaCollectDate{
    entity.NewAreaCollectDate(*kinds["kanen"], time.Date(2023, 6, 2, 0, 0, 0, 0, loc).Format("2006-01-02"), *area),
    entity.NewAreaCollectDate(*kinds["kanen"], time.Date(2023, 6, 6, 0, 0, 0, 0, loc).Format("2006-01-02"), *area),
    entity.NewAreaCollectDate(*kinds["kanen"], time.Date(2023, 6, 9, 0, 0, 0, 0, loc).Format("2006-01-02"), *area),
    entity.NewAreaCollectDate(*kinds["kanen"], time.Date(2023, 6, 13, 0, 0, 0, 0, loc).Format("2006-01-02"), *area),
    entity.NewAreaCollectDate(*kinds["kanen"], time.Date(2023, 6, 16, 0, 0, 0, 0, loc).Format("2006-01-02"), *area),
    entity.NewAreaCollectDate(*kinds["kanen"], time.Date(2023, 6, 20, 0, 0, 0, 0, loc).Format("2006-01-02"), *area),
    entity.NewAreaCollectDate(*kinds["kanen"], time.Date(2023, 6, 23, 0, 0, 0, 0, loc).Format("2006-01-02"), *area),
    entity.NewAreaCollectDate(*kinds["kanen"], time.Date(2023, 6, 27, 0, 0, 0, 0, loc).Format("2006-01-02"), *area),
    entity.NewAreaCollectDate(*kinds["kanen"], time.Date(2023, 6, 30, 0, 0, 0, 0, loc).Format("2006-01-02"), *area),
    entity.NewAreaCollectDate(*kinds["kanen"], time.Date(2023, 7, 4, 0, 0, 0, 0, loc).Format("2006-01-02"), *area),
    entity.NewAreaCollectDate(*kinds["kanen"], time.Date(2023, 7, 7, 0, 0, 0, 0, loc).Format("2006-01-02"), *area),
    entity.NewAreaCollectDate(*kinds["kanen"], time.Date(2023, 7, 11, 0, 0, 0, 0, loc).Format("2006-01-02"), *area),
    entity.NewAreaCollectDate(*kinds["kanen"], time.Date(2023, 7, 14, 0, 0, 0, 0, loc).Format("2006-01-02"), *area),
    entity.NewAreaCollectDate(*kinds["kanen"], time.Date(2023, 7, 18, 0, 0, 0, 0, loc).Format("2006-01-02"), *area),
    entity.NewAreaCollectDate(*kinds["kanen"], time.Date(2023, 7, 21, 0, 0, 0, 0, loc).Format("2006-01-02"), *area),
    entity.NewAreaCollectDate(*kinds["kanen"], time.Date(2023, 7, 25, 0, 0, 0, 0, loc).Format("2006-01-02"), *area),
    entity.NewAreaCollectDate(*kinds["kanen"], time.Date(2023, 7, 28, 0, 0, 0, 0, loc).Format("2006-01-02"), *area),
    entity.NewAreaCollectDate(*kinds["kanen"], time.Date(2023, 8, 1, 0, 0, 0, 0, loc).Format("2006-01-02"), *area),
    entity.NewAreaCollectDate(*kinds["kanen"], time.Date(2023, 8, 4, 0, 0, 0, 0, loc).Format("2006-01-02"), *area),
    entity.NewAreaCollectDate(*kinds["kanen"], time.Date(2023, 8, 8, 0, 0, 0, 0, loc).Format("2006-01-02"), *area),
    entity.NewAreaCollectDate(*kinds["kanen"], time.Date(2023, 8, 11, 0, 0, 0, 0, loc).Format("2006-01-02"), *area),
    entity.NewAreaCollectDate(*kinds["kanen"], time.Date(2023, 8, 15, 0, 0, 0, 0, loc).Format("2006-01-02"), *area),
    entity.NewAreaCollectDate(*kinds["kanen"], time.Date(2023, 8, 18, 0, 0, 0, 0, loc).Format("2006-01-02"), *area),
    entity.NewAreaCollectDate(*kinds["kanen"], time.Date(2023, 8, 22, 0, 0, 0, 0, loc).Format("2006-01-02"), *area),
    entity.NewAreaCollectDate(*kinds["kanen"], time.Date(2023, 8, 25, 0, 0, 0, 0, loc).Format("2006-01-02"), *area),
    entity.NewAreaCollectDate(*kinds["kanen"], time.Date(2023, 8, 29, 0, 0, 0, 0, loc).Format("2006-01-02"), *area),
    entity.NewAreaCollectDate(*kinds["funen"], time.Date(2023, 6, 1, 0, 0, 0, 0, loc).Format("2006-01-02"), *area),
    entity.NewAreaCollectDate(*kinds["funen"], time.Date(2023, 6, 15, 0, 0, 0, 0, loc).Format("2006-01-02"), *area),
    entity.NewAreaCollectDate(*kinds["funen"], time.Date(2023, 7, 6, 0, 0, 0, 0, loc).Format("2006-01-02"), *area),
    entity.NewAreaCollectDate(*kinds["funen"], time.Date(2023, 7, 20, 0, 0, 0, 0, loc).Format("2006-01-02"), *area),
    entity.NewAreaCollectDate(*kinds["funen"], time.Date(2023, 8, 3, 0, 0, 0, 0, loc).Format("2006-01-02"), *area),
    entity.NewAreaCollectDate(*kinds["funen"], time.Date(2023, 8, 17, 0, 0, 0, 0, loc).Format("2006-01-02"), *area),
    entity.NewAreaCollectDate(*kinds["shigen"], time.Date(2023, 6, 7, 0, 0, 0, 0, loc).Format("2006-01-02"), *area),
    entity.NewAreaCollectDate(*kinds["shigen"], time.Date(2023, 6, 14, 0, 0, 0, 0, loc).Format("2006-01-02"), *area),
    entity.NewAreaCollectDate(*kinds["shigen"], time.Date(2023, 6, 21, 0, 0, 0, 0, loc).Format("2006-01-02"), *area),
    entity.NewAreaCollectDate(*kinds["shigen"], time.Date(2023, 6, 28, 0, 0, 0, 0, loc).Format("2006-01-02"), *area),
    entity.NewAreaCollectDate(*kinds["shigen"], time.Date(2023, 7, 5, 0, 0, 0, 0, loc).Format("2006-01-02"), *area),
    entity.NewAreaCollectDate(*kinds["shigen"], time.Date(2023, 7, 12, 0, 0, 0, 0, loc).Format("2006-01-02"), *area),
    entity.NewAreaCollectDate(*kinds["shigen"], time.Date(2023, 7, 19, 0, 0, 0, 0, loc).Format("2006-01-02"), *area),
    entity.NewAreaCollectDate(*kinds["shigen"], time.Date(2023, 7, 26, 0, 0, 0, 0, loc).Format("2006-01-02"), *area),
    entity.NewAreaCollectDate(*kinds["shigen"], time.Date(2023, 8, 2, 0, 0, 0, 0, loc).Format("2006-01-02"), *area),
    entity.NewAreaCollectDate(*kinds["shigen"], time.Date(2023, 8, 9, 0, 0, 0, 0, loc).Format("2006-01-02"), *area),
    entity.NewAreaCollectDate(*kinds["shigen"], time.Date(2023, 8, 16, 0, 0, 0, 0, loc).Format("2006-01-02"), *area),
    entity.NewAreaCollectDate(*kinds["shigen"], time.Date(2023, 8, 23, 0, 0, 0, 0, loc).Format("2006-01-02"), *area),
    entity.NewAreaCollectDate(*kinds["shigen"], time.Date(2023, 8, 30, 0, 0, 0, 0, loc).Format("2006-01-02"), *area),
 }

  mockAreaCollectWeekdayRepository.EXPECT().FindByAreaId(area.ID).Return(mockAreaCollectWeekdays)

  s := NewAreaCollectWeekdayService(mockAreaCollectWeekdayRepository)
  result := s.ConvertByAreaWithAroundMonths(area, year, month)

  t.Run("[正常系] 指定した年月の前月と来月を含めた日付の配列を返す", func(t *testing.T) {
    assert.ElementsMatch(t, expected, result)
  })
}
