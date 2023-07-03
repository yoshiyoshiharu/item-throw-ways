package entity

import "time"

type AreaCollectionDate struct {
  Id  int
  Item Item
  Kind Kind
  Date time.Time
}
