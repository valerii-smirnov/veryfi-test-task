package domain

import "github.com/veryfi/veryfi-go/veryfi/scheme"

type Document struct {
	ID                 uint
	FileName           string
	VeryfiDocumentID   uint
	VeryfiDocumentInfo *scheme.Document
}
