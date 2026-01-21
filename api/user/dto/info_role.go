package dto

import (
	"github.com/afteracademy/goserve-example-api-server-mongo/api/user/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InfoRole struct {
	ID   primitive.ObjectID `json:"_id" binding:"required" validate:"required"`
	Code model.RoleCode     `json:"code" binding:"required" validate:"required,uppercase"`
}

func NewInfoRole(role *model.Role) *InfoRole {
	return &InfoRole{
		ID:   role.ID,
		Code: role.Code,
	}
}
