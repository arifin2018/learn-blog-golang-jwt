package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type (
	User struct{
		ID	int `json:"id" gorm:"primary_key"`
		Email string `json:"email" gorm:"not null;index:unique;unique"`
		Nickname string `json:"nickname" gorm:"not null"`
		Password string `json:"password" gorm:"not null"`
		Token string `json:"-" gorm:"type:text"`
	}
)

func (user *User) SaveUser(db *gorm.DB) (*User, error) {
	//turn password into hash
    hashedPassword, errPassword := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if errPassword != nil {
        return &User{}, errPassword
    }
	user.Password = string(hashedPassword)

	var err error = db.Create(&user).Error
	if err != nil {
        return &User{}, err
    }
    return user, nil
}