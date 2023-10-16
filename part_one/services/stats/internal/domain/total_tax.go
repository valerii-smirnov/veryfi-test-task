package domain

type TotalTax struct {
	Total float64
}

type TotalDiscount struct {
	Total float64
}

type GeographyInfo struct {
	Address    string
	Lat        float64
	Lng        float64
	TotalPrice float64
}
