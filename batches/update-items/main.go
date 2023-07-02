package main

import (
	"encoding/csv"
	"fmt"
	"net/http"

	"golang.org/x/text/transform"

	"golang.org/x/text/encoding/japanese"
)

const (
  API_URL = "https://www.city.bunkyo.lg.jp/library/opendata-bunkyo/01tetsuduki-kurashi/06bunbetuhinmoku/bunbetuhinmoku.csv"
)

func main() {
  updateItemsFromCsv()
}

func updateItemsFromCsv() error {
  resp, err := http.Get(
    API_URL,
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

  for _, row := range rows {
    fmt.Println(row)
  }

  return nil
}
