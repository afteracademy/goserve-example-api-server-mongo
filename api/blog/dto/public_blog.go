package dto

import (
	"time"

	"github.com/afteracademy/goserve-example-api-server-mongo/api/blog/model"
	userModel "github.com/afteracademy/goserve-example-api-server-mongo/api/user/model"
	"github.com/afteracademy/goserve/v2/utility"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PublicBlog struct {
	ID          primitive.ObjectID `json:"_id" binding:"required" validate:"required"`
	Title       string             `json:"title" validate:"required,min=3,max=500"`
	Description string             `json:"description" validate:"required,min=3,max=2000"`
	Text        string             `json:"text" validate:"required,max=50000"`
	Slug        string             `json:"slug" validate:"required,min=3,max=200"`
	Author      *InfoAuthor        `json:"author,omitempty" validate:"required,omitempty"`
	ImgURL      *string            `json:"imgUrl,omitempty" validate:"omitempty,uri,max=200"`
	Score       *float64           `json:"score,omitempty" validate:"omitempty,min=0,max=1"`
	Tags        *[]string          `json:"tags,omitempty" validate:"omitempty,dive,uppercase"`
	PublishedAt *time.Time         `json:"publishedAt,omitempty"`
}

func EmptyInfoPublicBlog() *PublicBlog {
	return &PublicBlog{}
}

func NewPublicBlog(blog *model.Blog, author *userModel.User) (*PublicBlog, error) {
	b, err := utility.MapTo[PublicBlog](blog)
	if err != nil {
		return nil, err
	}

	b.Author, err = utility.MapTo[InfoAuthor](author)
	if err != nil {
		return nil, err
	}

	return b, err
}

func (d *PublicBlog) GetValue() *PublicBlog {
	return d
}

func (b *PublicBlog) ValidateErrors(errs validator.ValidationErrors) ([]string, error) {
	return utility.FormatValidationErrors(errs), nil
}
