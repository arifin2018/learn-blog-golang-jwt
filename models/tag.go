package models

type (
	Tag struct{
		ID	int `json:"id" gorm:"primary_key"`
		Name string `json:"name" form:"name" binding:"required" gorm:"index:idx_name,unique"`
	}

	TagPost struct{
		ID	int `json:"id" gorm:"primary_key"`
		Name string `json:"name"`
		Post []GetPost `json:"post" gorm:"many2many:Post_Tags;foreignKey:ID;joinForeignKey:tag_id;References:ID;joinReferences:post_id"`
	}
)

func (TagPost) TableName() string {
	return "tags"
}