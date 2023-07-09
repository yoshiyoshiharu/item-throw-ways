package entity

import (
	"time"

	"gorm.io/gorm"
)

type AreaCollectionDate struct {
  gorm.Model
  Id  int `json:"id"`
  ItemId int `json:"item_id"`
  KindId int `json:"kind_id"`
  Date time.Time `json:"date"`
}
