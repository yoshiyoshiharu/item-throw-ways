package entity

import (
	"gorm.io/gorm"
)

type AreaCollectDate struct {
  gorm.Model
  Id  int `json:"id"`
  AreaId int `json:"area_id"`
  Area Area `json:"area"`
  KindId int `json:"kind_id"`
  Kind Kind `json:"kind"`
  Weekday int `json:"weekday"`
  Lap int `json:"lap"`
}
