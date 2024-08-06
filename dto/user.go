package dto

type User struct {
	Email    string `gorm:"column:email" json:"email"`
	Name     string `gorm:"column:name" json:"name"`
	UserType string `gorm:"column:user_type" json:"user_type"`
}

func (User) TableName() string {
	return "user_tab"
}
