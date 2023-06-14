package models

type (
	Image struct{
		ID	int `json:"id" gorm:"primary_key"`
		Image_url string `json:"image_url" form:"image_url" gorm:"type:text" binding:"required"`
	}
)