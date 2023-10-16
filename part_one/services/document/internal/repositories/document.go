package repositories

import (
	"context"
	"encoding/json"
	"github.com/veryfi/veryfi-go/veryfi/scheme"

	"github.com/valerii-smirnov/veryfi-test-task/part_one/document/internal/domain"
	"github.com/valerii-smirnov/veryfi-test-task/part_one/document/internal/repositories/models"

	"gorm.io/gorm"
)

type Document struct {
	db *gorm.DB
}

func NewFile(db *gorm.DB) *Document {
	return &Document{db: db}
}

func (d Document) Create(ctx context.Context, document domain.Document) (domain.Document, error) {
	jsonDocument, err := json.Marshal(document.VeryfiDocumentInfo)
	if err != nil {
		return domain.Document{}, err
	}

	modelFile := models.Document{
		FileName:           document.FileName,
		VeryfiDocumentID:   document.VeryfiDocumentID,
		VeryfiDocumentInfo: jsonDocument,
	}

	if err := d.db.WithContext(ctx).Create(&modelFile).Error; err != nil {
		return domain.Document{}, err
	}

	document.ID = modelFile.ID

	return document, nil
}

func (d Document) Get(ctx context.Context, id uint) (domain.Document, error) {
	var docModel models.Document
	if err := d.db.WithContext(ctx).
		Where("deleted_at IS NULL").
		Last(&docModel, id).Error; err != nil {
		return domain.Document{}, err
	}

	var docInfo scheme.Document
	if err := json.Unmarshal(docModel.VeryfiDocumentInfo, &docInfo); err != nil {
		return domain.Document{}, err
	}

	return domain.Document{
		ID:                 docModel.ID,
		FileName:           docModel.FileName,
		VeryfiDocumentID:   docModel.VeryfiDocumentID,
		VeryfiDocumentInfo: &docInfo,
	}, nil
}

func (d Document) GetByFileName(ctx context.Context, fileName string) (domain.Document, error) {
	var docModel models.Document
	if err := d.db.WithContext(ctx).
		Where("deleted_at IS NULL").
		Last(&docModel, "file_name = ?", fileName).Error; err != nil {
		return domain.Document{}, err
	}

	var docInfo scheme.Document
	if err := json.Unmarshal(docModel.VeryfiDocumentInfo, &docInfo); err != nil {
		return domain.Document{}, err
	}

	return domain.Document{
		ID:                 docModel.ID,
		FileName:           docModel.FileName,
		VeryfiDocumentID:   docModel.VeryfiDocumentID,
		VeryfiDocumentInfo: &docInfo,
	}, nil

}

func (d Document) Delete(ctx context.Context, id uint) error {
	if err := d.db.Delete(&models.Document{}, id).Error; err != nil {
		return err
	}

	return nil
}
