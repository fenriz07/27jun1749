package models

import "gorm.io/gorm"

type Counter struct {
	gorm.Model
	Code string `gorm:"column:code;type:varchar(20)"`
}

func (Counter) TableName() string {
	return "counter"
}
