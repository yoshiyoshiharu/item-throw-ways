package entity

import (
	"time"
)

type AreaCollectWeekday struct {
	ID      int  `json:"id"`
	AreaId  int  `json:"area_id"`
	Area    Area `json:"area"`
	KindId  int  `json:"kind_id"`
	Kind    Kind `json:"kind"`
	Weekday time.Weekday  `json:"weekday"`
	Lap     int  `json:"lap"`
}
