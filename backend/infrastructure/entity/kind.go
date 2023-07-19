package entity

type Kind struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func NewKind(id int, name string) *Kind {
	return &Kind{
		ID:   id,
		Name: name,
	}
}
