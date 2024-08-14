package dto

type Company struct {
	ID          int       `gorm:"column:id" json:"id"`
	Name        string    `gorm:"column:name" json:"name"`
	PlaceId     string    `gorm:"place_id" json:"place_id"`
	Location    string    `gorm:"location" json:"location"`
	Description string    `gorm:"description" json:"description"`
	Service     []Service `gorm:"-" json:"service"`
}

func (Company) TableName() string {
	return "company_tab"
}
