package models

type (
	User struct{
		ID	int `json:"id" gorm:"primary_key"`
		Email string `json:"email" gorm:"not null;index"`
		Nickname string `json:"nickname" gorm:"not null"`
		Password string `json:"password" gorm:"not null"`
		Token string `json:"token" gorm:"type:text"`
	}
)