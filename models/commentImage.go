package models

type (
	CommentImage struct{
		CommentId int `json:"comment_id" form:"comment_id"`
		ImageId int `json:"image_id" form:"image_id"`
	}
)