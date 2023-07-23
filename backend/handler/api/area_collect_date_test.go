package handler

import (
	"testing"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/yoshiyoshiharu/item-throw-ways/infrastructure/entity"
	mock_service "github.com/yoshiyoshiharu/item-throw-ways/mock/domain/service/api"
)

func TestAreaCollectDateHandler_FindAll (t *testing.T) {
  ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAreaCollectWeekdayService := mock_service.NewMockAreaCollectWeekdayService(ctrl)
	handler := NewAreaCollectDateHandler(mockAreaCollectWeekdayService)

  areas := []*entity.Area{
    entity.NewArea(1, "Area 1"),
  }
  kinds := []*entity.Kind{
    entity.NewKind(1, "可燃ごみ"),
    entity.NewKind(2, "不燃ごみ"),
    entity.NewKind(3, "資源"),
    entity.NewKind(4, "粗大ごみ"),
  }
  areaCollectDates := []*entity.AreaCollectDate{
    entity.NewAreaCollectDate(*kinds[0], "2023-07-01", *areas[0]),
    entity.NewAreaCollectDate(*kinds[1], "2023-07-02", *areas[0]),
    entity.NewAreaCollectDate(*kinds[1], "2023-07-03", *areas[0]),
  }

  mockAreaCollectWeekdayService.EXPECT().ConvertByAreaWithAroundMonths(1, 2023, time.Month(7)).Return(areaCollectDates)

  request := events.APIGatewayProxyRequest{
    QueryStringParameters: map[string]string{
      "area_id": "1",
      "year": "2023",
      "month": "7",
    },
  }
  resp, err := handler.FindAll(request)

  t.Run("[正常系] エリアの回収日一覧をJSONで返すこと", func(t *testing.T) {
    assert.NoError(t, err)

    assert.Equal(t, 200, resp.StatusCode)
    assert.Equal(t, `[{"kind":{"id":1,"name":"可燃ごみ"},"date":"2023-07-01"},{"kind":{"id":2,"name":"不燃ごみ"},"date":"2023-07-02"},{"kind":{"id":2,"name":"不燃ごみ"},"date":"2023-07-03"}]`, resp.Body)
  })
}
