package dto

import "github.com/google/uuid"

type CreateSideRequest struct {
	UserID uuid.UUID
	Name string
}

type CreateSideResponse struct {
	ID uuid.UUID
}