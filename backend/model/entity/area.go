package entity

import "gorm.io/gorm"

type Area struct {
  gorm.Model
  Id  int `json:"id"`
  Name string `json:"name"`
}
