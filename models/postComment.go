package models

type(
	PostComment struct{
		PostId	int `json:"post_id" form:"post_id"`
		CommentId	int `json:"comment_id" form:"comment_id"`
	}
)