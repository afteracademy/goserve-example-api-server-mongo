package dto

import (
	"github.com/afteracademy/goserve-example-api-server-mongo/api/user/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InfoPublicUser struct {
	ID            primitive.ObjectID `json:"_id" binding:"required" validate:"required"`
	Name          string             `json:"name" binding:"required" validate:"required"`
	ProfilePicURL *string            `json:"profilePicUrl,omitempty" validate:"omitempty,url"`
}

func NewInfoPublicUser(user *model.User) *InfoPublicUser {
	roles := make([]*InfoRole, len(user.Roles))
	for i, role := range user.RoleDocs {
		roles[i] = NewInfoRole(role)
	}

	return &InfoPublicUser{
		ID:            user.ID,
		Name:          user.Name,
		ProfilePicURL: user.ProfilePicURL,
	}
}
