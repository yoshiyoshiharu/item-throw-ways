package entity

import "time"

type AreaCollectDate struct {
  Kind string `json:"kind"`
  Date time.Time `json:"date"`
}
