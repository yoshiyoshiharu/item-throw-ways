package entity

import "time"

type AreaCollectionDate struct {
  Id  int `json:"id"`
  ItemId int `json:"item_id"`
  KindId int `json:"kind_id"`
  Date time.Time `json:"date"`
}
