package dto

import "github.com/google/uuid"

type JoinSideRequest struct {
	SideID uuid.UUID
	UserID uuid.UUID
}

type JoinSideResponse struct {
	MembershipID uuid.UUID	`json:"membership_id"`
}