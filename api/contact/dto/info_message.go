package dto

import (
	"time"

	"github.com/afteracademy/goserve/v2/utility"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InfoMessage struct {
	ID        primitive.ObjectID `json:"_id" binding:"required"`
	Type      string             `json:"type" binding:"required"`
	Msg       string             `json:"msg" binding:"required"`
	CreatedAt time.Time          `json:"createdAt" binding:"required"`
}

func EmptyInfoMessage() *InfoMessage {
	return &InfoMessage{}
}

func (d *InfoMessage) GetValue() *InfoMessage {
	return d
}

func (d *InfoMessage) ValidateErrors(errs validator.ValidationErrors) ([]string, error) {
	return utility.FormatValidationErrors(errs), nil
}
