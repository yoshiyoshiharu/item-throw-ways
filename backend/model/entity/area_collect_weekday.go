package entity

import (
	"time"

	"gorm.io/gorm"
)

type AreaCollectWeekday struct {
	gorm.Model
	Id      int  `json:"id"`
	AreaId  int  `json:"area_id"`
	Area    Area `json:"area"`
	KindId  int  `json:"kind_id"`
	Kind    Kind `json:"kind"`
	Weekday time.Weekday  `json:"weekday"`
	Lap     int  `json:"lap"`
}
