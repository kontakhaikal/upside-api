package dto

import "github.com/google/uuid"

type CreateSideRequest struct {
	UserID uuid.UUID
	Name string	`json:"name"`
}

type CreateSideResponse struct {
	SideID uuid.UUID	`json:"side_id"`
}