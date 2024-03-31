package services

import (
	"errors"
	"gorm.io/gorm"
	"webdriver_server/models"
)

type FileService struct {
	db *gorm.DB
}

func NewFileService(db *gorm.DB) *FileService {
	return &FileService{
		db: db,
	}
}

func (service *FileService) CreateFile(file *models.Files) error {
	return service.db.Create(file).Error
}

func (service *FileService) GetFileByID(id uint) (*models.Files, error) {
	var file models.Files
	if err := service.db.Preload("Owner").First(&file, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &file, nil
}

func (service *FileService) UpdateFile(file *models.Files) error {
	return service.db.Save(file).Error
}

func (service *FileService) DeleteFile(id uint) error {
	return service.db.Delete(&models.Files{}, id).Error
}

func (service *FileService) GetAllFiles() ([]models.Files, error) {
	var files []models.Files
	if err := service.db.Preload("Owner").Find(&files).Error; err != nil {
		return nil, err
	}
	return files, nil
}
