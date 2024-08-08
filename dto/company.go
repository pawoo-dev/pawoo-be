package dto

type Company struct {
	ID   int    `gorm:"column:id"`
	Name string `gorm:"column:name"`
}

func (Company) TableName() string {
	return "company_tab"
}
