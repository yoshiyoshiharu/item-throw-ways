package entity

type Kind struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func NewKind(id int, name string, items []Item) *Kind {
	return &Kind{
		ID:   id,
		Name: name,
	}
}
