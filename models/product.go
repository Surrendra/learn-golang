package models

type Product struct {
	Id          int64  `gorm:"primaryKey"json:"id"`
	Name        string `gorm:"type:varchar(200)" json:"name"`
	Description string `gorm:"type:text -all" json:"description"`
}
