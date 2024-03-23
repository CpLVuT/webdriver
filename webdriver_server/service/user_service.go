package services

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"webdriver_server/models" // 请将yourproject替换为你的项目导入路径
)

type UserService struct {
	db *gorm.DB
}

// NewUserService 创建一个新的UserService实例
func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		db: db,
	}
}

// CreateUser 创建一个新用户
func (s *UserService) CreateUser(user *models.User) error {
	// 检查邮箱是否已经存在
	exists, err := s.EmailExists(user.Email)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("email already exists")
	}

	// 创建用户
	result := s.db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// EmailExists 检查邮箱是否已经存在
func (s *UserService) EmailExists(email string) (bool, error) {
	var count int64
	s.db.Model(&models.User{}).Where("email = ?", email).Count(&count)
	return count > 0, nil
}

// GetUserByEmail 通过邮箱获取用户
func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := s.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}
	return &user, nil
}

// UpdateUser 更新用户信息
func (s *UserService) UpdateUser(userID uint, updateData map[string]interface{}) error {
	if password, ok := updateData["password"].(string); ok && password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		updateData["password"] = string(hashedPassword)
	}

	result := s.db.Model(&models.User{}).Where("id = ?", userID).Updates(updateData)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteUserByID 通过ID删除用户
func (s *UserService) DeleteUserByID(userID uint) error {
	result := s.db.Delete(&models.User{}, userID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no user found")
	}
	return nil
}

func (s *UserService) PasswordCheck(user *models.User, password string) error {
	err := user.CheckPassword(password)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) ExistanceCheck(email string) (bool, error) {
	var user models.User
	tb := s.db.Model(&models.User{})

	// Check if exists
	result := tb.Where("email = ?", email).Find(&user)
	if result.Error != nil {
		return false, result.Error
	}
	if result.RowsAffected > 0 {
		return true, nil
	}
	return false, nil
}
