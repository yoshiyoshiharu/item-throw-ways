package entity

type Kind struct {
	ID int `json:"id"`
  Name string `json:"name"`
  Items []Item `json:"items" gorm:"many2many:item_kinds;"`
}

