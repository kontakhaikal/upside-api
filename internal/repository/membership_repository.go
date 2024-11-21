package repository

import "github.com/fkrhykal/upside-api/internal/entity"

type MemberShipRepository[T any] interface {
	Save(ctx Context[T], membership *entity.MemberShip) (*entity.MemberShip, error)
}