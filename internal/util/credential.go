package util

import "github.com/fkrhykal/upside-api/internal/dto"


type CredentialUtil interface {
	GenerateToken(credential *dto.UserCredential) (dto.CredentialToken, error)
	VerifyToken(token dto.CredentialToken) error
}