package models

type(
	PostImage struct{
		PostId	int `json:"post_id" form:"post_id"`
		ImageId	int `json:"image_id" form:"image_id"`
	}
)