package models

import (
	"gorm.io/gorm"
	"time"
)

const ReceiptTableName = "receipt"

type Receipt struct {
	gorm.Model
	DocumentID    uint
	Category      string
	InvoiceNumber string
	Currency      string
	Tax           float64
	Total         float64
	Date          time.Time
	Geography     *Geography
	Items         []Item
}

func (r Receipt) TableName() string {
	return ReceiptTableName
}
