package models

type Product struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"varchar(300)" json:"name"`
	Description string `gorm:"text" json:"description"`
}
