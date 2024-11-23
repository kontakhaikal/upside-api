package dto

import "github.com/google/uuid"

type CredentialToken string

type UserCredential struct {
	ID uuid.UUID
}
