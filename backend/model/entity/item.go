package entity

import "gorm.io/gorm"

type Item struct {
  gorm.Model
  Id      int    `json:"id"`
  Name    string `json:"name"`
  NameKana string `json:"name_kana"`
  Price   int    `json:"price"`
  Remarks string `json:"remarks"`
  Kinds  []Kind `json:"kinds" gorm:"many2many:item_kinds;"`
}
