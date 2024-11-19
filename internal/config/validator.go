package config

import (
	"github.com/fkrhykal/upside-api/internal/util"
	"github.com/go-playground/validator/v10"
)


type PlaygroundValidator struct {
	validator *validator.Validate
}

func (p *PlaygroundValidator) ValidateDTO(dto any) error {
	return p.validator.Struct(dto)
}

func NewPlaygoundValidator() util.Validator {
	return &PlaygroundValidator{
		validator: validator.New(),
	}
}