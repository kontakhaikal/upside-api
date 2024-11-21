package entity

import "github.com/google/uuid"

type Role string

var Author Role = "author"

var Admin Role = "admin"

var Member Role = "member"

type MemberShip struct {
	ID uuid.UUID
	UserID uuid.UUID
	Side uuid.UUID
	Role
}