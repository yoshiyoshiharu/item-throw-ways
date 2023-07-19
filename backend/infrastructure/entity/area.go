package entity

type Area struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func NewArea(id int, name string) *Area {
	return &Area{
		ID:   id,
		Name: name,
	}
}
