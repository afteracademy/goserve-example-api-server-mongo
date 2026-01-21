package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateBlog struct {
	ID          primitive.ObjectID `json:"_id" binding:"required" validate:"required"`
	Title       *string            `json:"title" validate:"omitempty,min=3,max=500"`
	Description *string            `json:"description" validate:"omitempty,min=3,max=2000"`
	DraftText   *string            `json:"draftText" validate:"omitempty,max=50000"`
	Slug        *string            `json:"slug" validate:"omitempty,min=3,max=200"`
	ImgURL      *string            `json:"imgUrl" validate:"omitempty,uri,max=200"`
	Tags        *[]string          `json:"tags" validate:"omitempty,min=1,dive,uppercase"`
}
