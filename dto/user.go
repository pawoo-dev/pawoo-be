package dto

type User struct {
	ID          int    `gorm:"column:id"`
	Email       string `gorm:"column:email" json:"email"`
	Name        string `gorm:"column:name" json:"name"`
	UserType    string `gorm:"column:user_type" json:"user_type"`
	CompanyId   int    `gorm:"column:company_id"`
	CompanyName string `json:"company_name" gorm:"-"`
}

func (User) TableName() string {
	return "user_tab"
}
