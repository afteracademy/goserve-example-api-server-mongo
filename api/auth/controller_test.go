package auth

import (
	"net/http"
	"testing"

	"github.com/afteracademy/goserve-example-api-server-mongo/api/auth/dto"
	userDto "github.com/afteracademy/goserve-example-api-server-mongo/api/user/dto"
	"github.com/afteracademy/goserve-example-api-server-mongo/api/user/model"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/afteracademy/goserve/v2/network"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthController_SignupBadRequest(t *testing.T) {
	mockAuthProvider := new(network.MockAuthenticationProvider)
	mockAuthProvider.On("Middleware").Return(gin.HandlerFunc(func(ctx *gin.Context) {
		ctx.Next()
	}))

	mockAuthzProvider := new(network.MockAuthorizationProvider)
	mockAuthzProvider.On("Middleware", "ROLE").Return(gin.HandlerFunc(func(ctx *gin.Context) {
		ctx.Next()
	}))

	authService := new(MockService)

	c := NewController(mockAuthProvider, mockAuthzProvider, authService)

	rr := network.MockTestController(t, "POST", "/auth/signup/basic", `{}`, c)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Contains(t, rr.Body.String(), `"message":"email is required, password is required, name is required"`)
}

func TestAuthController_SignupSuccess(t *testing.T) {
	mockAuthProvider := new(network.MockAuthenticationProvider)
	mockAuthProvider.On("Middleware").Return(gin.HandlerFunc(func(ctx *gin.Context) {
		ctx.Next()
	}))

	mockAuthzProvider := new(network.MockAuthorizationProvider)
	mockAuthzProvider.On("Middleware", "ROLE").Return(gin.HandlerFunc(func(ctx *gin.Context) {
		ctx.Next()
	}))

	body := `{"email":"test@abc.com","password":"123456","name":"test name"}`

	singUpDto := &dto.SignUpBasic{
		Email:    "test@abc.com",
		Password: "123456",
		Name:     "test name",
	}

	authDto := &dto.UserAuth{
		User: &userDto.InfoPrivateUser{
			Name:  "test name",
			Email: "test@abc.com",
			ID:    primitive.NewObjectID(),
			Roles: []*userDto.InfoRole{
				{
					ID:   primitive.NewObjectID(),
					Code: model.RoleCodeLearner,
				},
			},
			ProfilePicURL: nil,
		},
		Tokens: &dto.UserTokens{
			AccessToken:  "access-token",
			RefreshToken: "refresh-token",
		},
	}

	authService := new(MockService)
	authService.On("SignUpBasic", singUpDto).Return(authDto, nil)

	c := NewController(mockAuthProvider, mockAuthzProvider, authService)

	rr := network.MockTestController(t, "POST", "/auth/signup/basic", body, c)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), `"message":"success"`)
}
