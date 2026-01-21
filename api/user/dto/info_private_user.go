package dto

import (
	"github.com/afteracademy/goserve-example-api-server-mongo/api/user/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InfoPrivateUser struct {
	ID            primitive.ObjectID `json:"_id" binding:"required" validate:"required"`
	Email         string             `json:"email" binding:"required" validate:"required,email"`
	Name          string             `json:"name" binding:"required" validate:"required"`
	ProfilePicURL *string            `json:"profilePicUrl,omitempty" validate:"omitempty,url"`
	Roles         []*InfoRole        `json:"roles" validate:"required,dive,required"`
}

func NewInfoPrivateUser(user *model.User) *InfoPrivateUser {
	roles := make([]*InfoRole, len(user.Roles))
	for i, role := range user.RoleDocs {
		roles[i] = NewInfoRole(role)
	}

	return &InfoPrivateUser{
		ID:            user.ID,
		Email:         user.Email,
		Name:          user.Name,
		ProfilePicURL: user.ProfilePicURL,
		Roles:         roles,
	}
}
