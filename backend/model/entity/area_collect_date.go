package entity

type AreaCollectDate struct {
	Kind   Kind   `json:"kind"`
	Date   string `json:"date"`
	Area   Area
	AreaID int
}

func NewAreaCollectDate(kind Kind, date string, area Area) *AreaCollectDate {
	return &AreaCollectDate{
		Kind:   kind,
		Date:   date,
		Area:   area,
		AreaID: area.ID,
	}
}
