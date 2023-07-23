package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	mock_service "github.com/yoshiyoshiharu/item-throw-ways/mock/domain/service/batch"
)

func TestAreaCollectWeekdayBatchHandler_UpdateAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAreaCollectWeekdayBatchService := mock_service.NewMockAreaCollectWeekdayBatchService(ctrl)
	handler := NewAreaCollectWeekdayBatchHandler(mockAreaCollectWeekdayBatchService)

	t.Run("[正常系] バッチ処理に成功したとき、成功したことをスラック通知すること", func(t *testing.T) {
		slackMockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				t.Errorf("expected POST method, got %s", r.Method)
			}

			var requestBody map[string]string
			err := json.NewDecoder(r.Body).Decode(&requestBody)
			if err != nil {
				t.Fatal(err)
			}

			expectedRequestBody := map[string]string{
				"text": "バッチ処理に成功しました。",
			}

			if requestBody["text"] != expectedRequestBody["text"] {
				t.Errorf("expected request body %v, got %v", expectedRequestBody, requestBody)
			}

			w.WriteHeader(http.StatusOK)
		}))

		defer slackMockServer.Close()
		slackWebhookUrl = slackMockServer.URL

		mockAreaCollectWeekdayBatchService.EXPECT().UpdateAll().Return(nil)
		handler.UpdateAll()
	})

	t.Run("[異常系] バッチ処理に失敗したとき、失敗したことをスラック通知すること", func(t *testing.T) {
		slackMockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				t.Errorf("expected POST method, got %s", r.Method)
			}

			var requestBody map[string]string
			err := json.NewDecoder(r.Body).Decode(&requestBody)
			if err != nil {
				t.Fatal(err)
			}

			expectedRequestBody := map[string]string{
				"text": "バッチ処理でエラーが発生しました。: エラー内容",
			}

			if requestBody["text"] != expectedRequestBody["text"] {
				t.Errorf("expected request body %v, got %v", expectedRequestBody, requestBody)
			}

			w.WriteHeader(http.StatusOK)
		}))

		defer slackMockServer.Close()
		slackWebhookUrl = slackMockServer.URL

		mockAreaCollectWeekdayBatchService.EXPECT().UpdateAll().Return(errors.New("エラー内容"))
		handler.UpdateAll()
	})
}
