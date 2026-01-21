package dto

import (
	"github.com/afteracademy/goserve-example-api-server-mongo/api/user/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InfoAuthor struct {
	ID            primitive.ObjectID `json:"_id" binding:"required" validate:"required"`
	Name          string             `json:"name" binding:"required" validate:"required"`
	ProfilePicURL *string            `json:"profilePicUrl,omitempty" validate:"omitempty,url"`
}

func NewInfoPrivateUser(user *model.User) *InfoAuthor {
	return &InfoAuthor{
		ID:            user.ID,
		Name:          user.Name,
		ProfilePicURL: user.ProfilePicURL,
	}
}
