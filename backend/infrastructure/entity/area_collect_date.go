package entity

type AreaCollectDate struct {
	Kind   Kind   `json:"kind"`
	Date   string `json:"date"`
	area   Area
	areaID int
}

func NewAreaCollectDate(kind Kind, date string, area Area) *AreaCollectDate {
	return &AreaCollectDate{
		Kind:   kind,
		Date:   date,
		area:   area,
		areaID: area.ID,
	}
}
