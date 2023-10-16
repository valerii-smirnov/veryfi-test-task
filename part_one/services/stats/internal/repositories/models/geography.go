package models

import "gorm.io/gorm"

const GeographyTableName = "receipt_geography"

type Geography struct {
	gorm.Model
	ReceiptID uint `gorm:"uniqueIndex"`
	Address   string
	Lat       float64
	Lng       float64
}

func (g Geography) TableName() string {
	return GeographyTableName
}
