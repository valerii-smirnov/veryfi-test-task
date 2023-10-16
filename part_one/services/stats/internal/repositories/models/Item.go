package models

import "gorm.io/gorm"

const ItemTableName = "receipt_item"

type Item struct {
	gorm.Model
	ReceiptID   uint
	Type        string
	Description string
	Discount    float64
	Quantity    float64
	Tax         float64
	TaxRate     float64
	Total       float64
}

func (i Item) TableName() string {
	return ItemTableName
}
