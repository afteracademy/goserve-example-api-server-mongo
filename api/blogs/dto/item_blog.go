package dto

import (
	"github.com/afteracademy/goserve-example-api-server-mongo/api/blog/model"
	"github.com/afteracademy/goserve-example-api-server-mongo/utils"
	"github.com/afteracademy/goserve/v2/utility"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ItemBlog struct {
	ID          primitive.ObjectID `json:"_id" binding:"required" validate:"required"`
	Title       string             `json:"title" validate:"required,min=3,max=500"`
	Description string             `json:"description" validate:"required,min=3,max=2000"`
	Slug        string             `json:"slug" validate:"required,min=3,max=200"`
	ImgURL      *string            `json:"imgUrl,omitempty" validate:"omitempty,uri,max=200"`
	Score       float64            `json:"score," validate:"required,min=0,max=1"`
	Tags        []string           `json:"tags" validate:"required,dive,uppercase"`
}

func NewItemBlog(blog *model.Blog) (*ItemBlog, error) {
	return utils.MapTo[ItemBlog](blog)
}

func EmptyItemBlog() *ItemBlog {
	return &ItemBlog{}
}

func (d *ItemBlog) GetValue() *ItemBlog {
	return d
}

func (b *ItemBlog) ValidateErrors(errs validator.ValidationErrors) ([]string, error) {
	return utility.FormatValidationErrors(errs), nil
}
