package dto

type Service struct {
	ID          int              `gorm:"column:id"`
	CompanyId   int              `gorm:"column:company_id"`
	Name        string           `gorm:"column:name" json:"name"`
	Description string           `gorm:"column:description" json:"description"`
	MultiSelect bool             `gorm:"column:multiselect" json:"multiselect"`
	Options     []ServiceOptions `json:"options" gorm:"-"`
}

func (Service) TableName() string {
	return "service_tab"
}

type ServiceOptions struct {
	ID          int    `gorm:"column:id"`
	ServiceId   int    `gorm:"column:service_id"`
	Name        string `gorm:"column:name" json:"name"`
	Description string `gorm:"column:description" json:"description"`
	Price       int    `gorm:"column:price" json:"price"`
	Duration    int    `gorm:"column:duration" json:"duration"`
}

func (ServiceOptions) TableName() string {
	return "service_option_tab"
}
