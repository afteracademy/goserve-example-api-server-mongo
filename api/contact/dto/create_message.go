package dto

import (
	"github.com/afteracademy/goserve/v2/utility"
	"github.com/go-playground/validator/v10"
)

type CreateMessage struct {
	Type string `json:"type" binding:"required,min=2,max=50"`
	Msg  string `json:"msg" binding:"required,min=0,max=2000"`
}

func EmptyCreateMessage() *CreateMessage {
	return &CreateMessage{}
}

func (d *CreateMessage) GetValue() *CreateMessage {
	return d
}

func (d *CreateMessage) ValidateErrors(errs validator.ValidationErrors) ([]string, error) {
	return utility.FormatValidationErrors(errs), nil
}
