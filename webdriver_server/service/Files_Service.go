package services

import (
	"errors"
	"gorm.io/gorm"
	"webdriver_server/models" // 引入你的模型包
)

// FileService 提供了对文件操作的接口定义
type FileService struct {
	db *gorm.DB
}

// NewFileService 创建一个新的FileService实例
func NewFileService(db *gorm.DB) *FileService {
	return &FileService{
		db: db,
	}
}

// CreateFile 创建一个新的文件记录
func (service *FileService) CreateFile(file *models.Files) error {
	return service.db.Create(file).Error
}

// GetFileByID 根据ID获取一个文件
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

// UpdateFile 更新一个文件的信息
func (service *FileService) UpdateFile(file *models.Files) error {
	return service.db.Save(file).Error
}

// DeleteFile 删除一个文件
func (service *FileService) DeleteFile(id uint) error {
	return service.db.Delete(&models.Files{}, id).Error
}

// GetAllFiles 获取所有文件
func (service *FileService) GetAllFiles() ([]models.Files, error) {
	var files []models.Files
	if err := service.db.Preload("Owner").Find(&files).Error; err != nil {
		return nil, err
	}
	return files, nil
}
