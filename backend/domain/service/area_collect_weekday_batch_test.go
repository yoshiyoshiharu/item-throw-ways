package service

import (
	"bytes"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	mock_repository "github.com/yoshiyoshiharu/item-throw-ways/mock/domain/repository"
)

func TestAreaCollectWeekdayBatch_UpdateAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAreaCollectWeekdayRepository := mock_repository.NewMockAreaCollectWeekdayRepository(ctrl)
	mockKindRepository := mock_repository.NewMockKindRepository(ctrl)

	areaCollectWeekdayService := NewAreaCollectWeekdayBatchService(mockAreaCollectWeekdayRepository, mockKindRepository)

	// モックAPIサーバーを使わずに、サンプルのCSVデータを直接使用する
	sampleCSVData := `エリア,通り,可燃ごみ,不燃ごみ,資源
地域A,通り1,月曜日・金曜日,第2・第4火曜日,水曜日
地域B,通り2,木曜日・土曜日,第1・第3火曜日,木曜日`
	csvMock := MockCSVReader([]byte(sampleCSVData))

	// リポジトリの期待される動作をセットアップ
	mockKindRepository.EXPECT().FindAll().Return(GetSampleKinds())

	// 削除と挿入の期待されるアイテムをセットアップ
	expectedAreaCollectWeekdays := GetExpectedAreaCollectWeekdays()
	mockAreaCollectWeekdayRepository.EXPECT().DeleteAndInsertAll(expectedAreaCollectWeekdays).Return(nil)

	// テスト対象のメソッドを実行
	err := areaCollectWeekdayService.UpdateAll(csvMock)

	// アサーションを行う
	assert.NoError(t, err) // エラーがないことを確認
}

// テスト用のCSVReaderを返すヘルパー関数
func MockCSVReader(data []byte) yourpackage.CSVReader {
	r := NewCSVReader(bytes.NewReader(data))
	return r
}
