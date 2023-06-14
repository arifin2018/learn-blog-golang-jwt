package models

type (
	Tag struct{
		ID	int `json:"id" gorm:"primary_key"`
		Name string `json:"name" form:"name" binding:"required"`
	}
)