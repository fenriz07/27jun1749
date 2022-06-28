package models

type Link struct {
	ID   string `gorm:"column:id;type:text;primaryKey"`
	Code string `gorm:"column:code;type:varchar(20);unique"`
}

func (Link) TableName() string {
	return "link"
}
