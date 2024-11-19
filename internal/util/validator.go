package util

type Validator interface {
	ValidateDTO(dto any) error
}