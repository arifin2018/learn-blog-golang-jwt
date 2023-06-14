package models

type(
	PostTag struct{
		PostId	int `json:"post_id" form:"post_id"`
		TagId	int `json:"tag_id" form:"tag_id"`
	}
)