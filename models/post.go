package models

import "time"

type (
	Post struct{
		ID	int `json:"id" form:"id" gorm:"primary_key"`
		Title string `json:"title" form:"title" gorm:"not null;type:varchar(200)" binding:"required"`
		Content string `json:"content" form:"content" gorm:"type:text" binding:"required"`
		User_id int `json:"user_id" form:"user_id" gorm:"not null"`
		Image_url []string `json:"image_url" form:"image_url" gorm:"-"`
		ImageId []int `json:"image_id" form:"image_id" gorm:"-"`
		TagId []int `json:"tag_id" form:"tag_id" gorm:"-"`
		CommentId []int `json:"comment_id" form:"comment_id" gorm:"-"`
		User PostUser `json:"user"`
		Image []Image `json:"image" gorm:"many2many:Post_Images;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
		Tag []Tag `json:"tag" gorm:"many2many:Post_Tags;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
		Comment []Comment `json:"comment" gorm:"many2many:Post_Comments;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
		Created_at time.Time `json:"created_at"`
	}

	PostReq struct{
		Title string `json:"title" form:"title" gorm:"not null;type:varchar(200)" binding:"required"`
		Content string `json:"content" form:"content" gorm:"type:text" binding:"required"`
		Image_url []string `json:"image_url" form:"image_url" gorm:"-"`
		TagId []int `json:"tag_id" form:"tag_id" gorm:"-"`
	}

	GetPost struct{
		ID	int `json:"id" form:"id" gorm:"primary_key"`
		Title string `json:"title" form:"title" gorm:"not null;type:varchar(200)" binding:"required"`
		Content string `json:"content" form:"content" gorm:"type:text" binding:"required"`
		User_id int `json:"-" gorm:"not null"`
		Image_url []string `json:"-" gorm:"-"`
		ImageId []int `json:"-" gorm:"-"`
		TagId []int `json:"-" gorm:"-"`
		CommentId []int `json:"-" gorm:"-"`
		User PostUser `json:"user"`
		Image []Image `json:"image" gorm:"many2many:Post_Images;foreignKey:ID;joinForeignKey:post_id;References:ID;joinReferences:image_id"`
		Tag []Tag `json:"tag" gorm:"many2many:Post_Tags;foreignKey:ID;joinForeignKey:post_id;References:ID;joinReferences:tag_id"`
		Comment []Comment `json:"comment" gorm:"many2many:Post_Comments;foreignKey:ID;joinForeignKey:post_id;References:ID;joinReferences:comment_id"`
		Created_at time.Time `json:"created_at"`
	}
)

func (GetPost) TableName() string {
	return "posts"
}