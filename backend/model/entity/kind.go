package entity

import "gorm.io/gorm"

type Kind struct {
  gorm.Model
	Id int `json:"id"`
  Name string `json:"name"`
  Items []Item `json:"items" gorm:"many2many:item_kinds;"`
}

