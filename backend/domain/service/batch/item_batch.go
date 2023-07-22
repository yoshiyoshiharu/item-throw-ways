package service

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/yoshiyoshiharu/item-throw-ways/domain/repository"
	"github.com/yoshiyoshiharu/item-throw-ways/infrastructure/entity"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

type ItemBatchService interface {
  UpdateAll() error
}

type itemBatchService struct {
  ir repository.ItemRepository
  kr repository.KindRepository
}

func NewItemBatchService(ir repository.ItemRepository, kr repository.KindRepository) *itemBatchService {
  return &itemBatchService{
    ir: ir,
    kr: kr,
  }
}

type RequestBody struct {
	AppId      string `json:"app_id"`
	OutputType string `json:"output_type"`
	Sentence   string `json:"sentence"`
}

type ResponseBody struct {
	Converted string `json:"converted"`
}
var (
  ItemsApiUrl                      = "https://www.city.bunkyo.lg.jp/library/opendata-bunkyo/01tetsuduki-kurashi/06bunbetuhinmoku/bunbetuhinmoku.csv"
  HiraganaTranslationApiUrl = "https://labs.goo.ne.jp/api/hiragana"
  CONCURRENCY                  = 10
)

func (s *itemBatchService) UpdateAll() error {
	var (
    wg    sync.WaitGroup
    mu    sync.Mutex
    items []entity.Item
  )

	allKinds := s.kr.FindAll()

	resp, err := http.Get(
		ItemsApiUrl,
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	r := csv.NewReader(transform.NewReader(resp.Body, japanese.ShiftJIS.NewDecoder()))
	rows, err := r.ReadAll()
	if err != nil {
		return err
	}

	itemChan := make(chan *entity.Item)
	semaphore := make(chan struct{}, CONCURRENCY)

	for i, row := range rows {
		// ヘッダー行はスキップ
		if i == 0 {
			continue
		}

		wg.Add(1)
		semaphore <- struct{}{} // セマフォに空きがでるまでブロック

		go func(insertId int, row []string) {
			itemId := insertId
			itemName := row[1]
			kindNames := GetKindsFromCell(row[2])
			price, _ := strconv.Atoi(row[3])
			remarks := row[4]

			itemNameKana, err := TranslateToHiragana(itemName)
			if err != nil {
        panic(err)
			}

			var kinds []entity.Kind
			for _, kindName := range kindNames {
				kind := findKind(kindName, allKinds)

				kinds = append(kinds, *kind)
			}

			itemChan <- entity.NewItem(
				itemId,
				itemName,
				itemNameKana,
				price,
				remarks,
				kinds,
			)

			<-semaphore
			wg.Done()
		}(i, row)

		go func() {
			for item := range itemChan {
				if itemExists(item.Name, items) {
					continue
				}
				mu.Lock()
				items = append(items, *item)
				mu.Unlock()
			}
		}()
	}

	wg.Wait()

  close(itemChan)

	s.ir.DeleteAndInsertAll(items)

  return nil
}

func GetKindsFromCell(str string) []string {
	return strings.Split(str, "、")
}

func TranslateToHiragana(name string) (string, error) {
	requestBody := &RequestBody{
		AppId:      os.Getenv("HIRAGANA_TRANSLATE_APP_ID"),
		OutputType: "hiragana",
		Sentence:   name,
	}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	req, _ := http.NewRequest("POST", HiraganaTranslationApiUrl, bytes.NewBuffer(jsonString))
	req.Header.Set("Content-Type", "application/json")
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	var responseBody ResponseBody
	json.NewDecoder(resp.Body).Decode(&responseBody)

	return responseBody.Converted, nil
}
