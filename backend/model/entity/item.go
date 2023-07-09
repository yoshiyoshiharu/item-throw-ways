package entity

import "gorm.io/gorm"

type Item struct {
  gorm.Model
  Id      int    `json:"id"`
  Name    string `json:"name"`
  Price   int    `json:"price"`
  Remarks string `json:"remarks"`
}
