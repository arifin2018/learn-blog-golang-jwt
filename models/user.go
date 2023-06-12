package models

import (
	"Blog/helpers"

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

	LoginUsers struct{
		Email string `json:"email"`
		Password string `json:"password"`
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

func VerifyPassword(password, hashedPassword string) error {
    return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (user *User) LoginCheck(email string, password string, db *gorm.DB) (string, User, error) {
    var err error

    u := User{}

    err = db.Model(User{}).Where("email = ?", email).Take(&u).Error

    if err != nil {
        return "",u, err
    }

    err = VerifyPassword(password, u.Password)

    if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
        return "",u, err
    }

    token, err := helpers.GenerateToken(uint(u.ID))

    if err != nil {
        return "", u,err
    }

    return token,u, nil
}