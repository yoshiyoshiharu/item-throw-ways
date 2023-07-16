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

func NewAreaCollectWeekday(area Area, kind Kind, weekday time.Weekday, lap int) *AreaCollectWeekday {
  return &AreaCollectWeekday{
    AreaId:  area.ID,
    Area:    area,
    KindId:  kind.ID,
    Kind:    kind,
    Weekday: weekday,
    Lap:     lap,
  }
}
