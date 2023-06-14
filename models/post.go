package models

import "time"

type (
	Post struct{
		ID	int `json:"id" form:"id" gorm:"primary_key"`
		Content string `json:"content" form:"content" gorm:"type:text" binding:"required"`
		User_id int `json:"user_id" form:"user_id" gorm:"not null"`
		Image_url []string `json:"image_url" form:"image_url" gorm:"-"`
		ImageId []int `json:"image_id" form:"image_id" gorm:"-"`
		TagId []int `json:"tag_id" form:"tag_id" gorm:"-"`
		CommentId []int `json:"comment_id" form:"comment_id" gorm:"-"`
		User PostUser `json:"user"`
		Image []Image `json:"image" gorm:"many2many:Post_Images"`
		Tag []Tag `json:"tag" gorm:"many2many:Post_Tags"`
		Comment []Comment `json:"comment" gorm:"many2many:Post_Comments"`
		Created_at time.Time `json:"created_at"`
	}
)