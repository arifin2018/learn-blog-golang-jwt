package models

import "time"

type (
	Comment struct{
		ID	int `json:"id" gorm:"primary_key"`
		Content string `json:"content" form:"content" gorm:"type:text" binding:"required"`
		User_id int `json:"user_id" form:"user_id" gorm:"not null"`
		User User `json:"user"`
		Created_at time.Time `json:"created_at"`
	}
)