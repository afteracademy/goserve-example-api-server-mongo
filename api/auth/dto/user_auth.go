package dto

import (
	"github.com/afteracademy/goserve-example-api-server-mongo/api/user/dto"
	"github.com/afteracademy/goserve-example-api-server-mongo/api/user/model"
)

type UserAuth struct {
	User   *dto.InfoPrivateUser `json:"user" validate:"required"`
	Tokens *UserTokens          `json:"tokens" validate:"required"`
}

func NewUserAuth(user *model.User, tokens *UserTokens) *UserAuth {
	return &UserAuth{
		User:   dto.NewInfoPrivateUser(user),
		Tokens: tokens,
	}
}
