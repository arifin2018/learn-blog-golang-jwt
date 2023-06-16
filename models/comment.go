package models

import "time"

type (
	Comment struct{
		ID	int `json:"id" gorm:"primary_key"`
		Content string `json:"content" form:"content" gorm:"type:text" binding:"required"`
		ImageUrl []string `json:"image_url" form:"image_url" gorm:"-"`
		User_id int `json:"user_id" form:"user_id" gorm:"not null"`
		User PostUser `json:"user" gorm:"foreignKey:User_id"`
		Image []Image `json:"image" gorm:"many2many:comment_images;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
		Created_at time.Time `json:"created_at"`
	}
	GetComment struct{
		ID	int `json:"id" gorm:"primary_key"`
		Content string `json:"content" form:"content"`
		User_id int `json:"-" form:"user_id"`
		User PostUser `json:"user" gorm:"foreignKey:User_id"`
		Image []Image `json:"image" gorm:"many2many:comment_images;foreignKey:ID;joinForeignKey:comment_id;References:ID;joinReferences:image_id"`
		Created_at time.Time `json:"created_at"`
	}
)
func (GetComment) TableName() string {
	return "comments"
}