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
		Nickname string `json:"nickname" gorm:"not null" binding:"required"`
		ImageUrl string `json:"imageUrl" gorm:"not null" binding:"required"`
		Password string `json:"password" gorm:"not null" binding:"required"`
		Token string `json:"-" gorm:"type:text"`
	}

	LoginUsers struct{
		Email string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

    RegisterUsers struct{
		Email string `json:"email"`
		Nickname string `json:"nickname" binding:"required"`
		ImageUrl string `json:"imageUrl" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

    LoginRes struct{
		Message string
        User    []UserToken
	}

    UserToken struct{
        Email string `json:"email"`
        Nickname string `json:"nickname"`
		ImageUrl string `json:"imageUrl" binding:"required"`
        Token string `json:"token"`
    }

    PostUser struct{
		ID	int `json:"id" gorm:"primary_key"`
		Email string `json:"email" gorm:"not null;index:unique;unique"`
		ImageUrl string `json:"imageUrl"`
		Nickname string `json:"nickname" gorm:"not null"`
	}
)

func (PostUser) TableName() string {
	return "users"
}

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