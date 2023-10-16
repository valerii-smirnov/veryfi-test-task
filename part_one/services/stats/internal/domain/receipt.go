package domain

import (
	"time"
)

type Receipt struct {
	ID            uint
	DocumentID    uint
	Category      string
	InvoiceNumber string
	Currency      string
	Tax           float64
	Total         float64
	Date          time.Time
	Geography     *Geography
	Items         []Item
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Geography struct {
	ID        uint
	Address   string
	Lat       float64
	Lng       float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Item struct {
	ID          uint
	Type        string
	Description string
	Discount    float64
	Quantity    float64
	Tax         float64
	TaxRate     float64
	Total       float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
