package entity

type Item struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	NameKana string `json:"name_kana"`
	Price    int    `json:"price"`
	Remarks  string `json:"remarks"`
	Kinds    []Kind `json:"kinds" gorm:"many2many:item_kinds;"`
}

func NewItem(id int, name string, nameKana string, price int, remarks string, kinds []Kind) *Item {
	return &Item{
		ID:       id,
		Name:     name,
		NameKana: nameKana,
		Price:    price,
		Remarks:  remarks,
		Kinds:    kinds,
	}
}
