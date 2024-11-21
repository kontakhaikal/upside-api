package repository

import "github.com/fkrhykal/upside-api/internal/entity"

type SideRepository[T any] interface {
	Save(ctx Context[T], side *entity.Side) (*entity.Side, error)
}