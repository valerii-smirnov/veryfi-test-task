package models

import "gorm.io/gorm"

const DocumentTableName = "document"

type Document struct {
	gorm.Model
	FileName           string
	VeryfiDocumentID   uint
	VeryfiDocumentInfo []byte
}

func (f Document) TableName() string {
	return DocumentTableName
}
