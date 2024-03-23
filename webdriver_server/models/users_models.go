package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	CommonModelFields
	Email    string  `json:"email" binding:"required" gorm:"uniqueIndex,size:64"`
	Mobile   *string `json:"mobile" gorm:"uniqueIndex,size:16"`
	Password string  `json:"-"`
	Nickname string  `json:"nickname"`
	Avatar   string  `json:"avatar"`
}

type AdminUser struct {
	CommonModelFields
	Email       string  `json:"email" binding:"required"`
	IsSuperUser bool    `json:"isSuperUser" binding:"required"`
	Mobile      *string `json:"mobile"`
	Password    string  `json:"-" binding:"required"`
	Nickname    string  `json:"nickname" binding:"required"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return
}

func (u *User) CheckPassword(pwd string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pwd))
	if err != nil {
		return err
	}
	return nil

}

func (u *AdminUser) BeforeCreate(tx *gorm.DB) (err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hash)
	return
}

func (u *AdminUser) CheckPassword(pwd string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pwd))
	if err != nil {
		return err
	}
	return nil
}

func UserMigrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&AdminUser{})
}
