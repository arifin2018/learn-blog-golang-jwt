package models

type (
	Post struct{
		ID	int `json:"id" gorm:"primary_key"`
		Content string `json:"content" gorm:"type:text" binding:"required"`
		User_id int `json:"user_id" gorm:"not null"`
		Image_id int `json:"image_id"`
		Tag_id int `json:"tag_id"`
		Comment_id int `json:"comment_id"`
		Created_at string `json:"created_at"`
	}
)