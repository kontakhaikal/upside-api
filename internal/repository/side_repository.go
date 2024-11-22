package repository

import (
	"errors"

	"github.com/fkrhykal/upside-api/internal/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SideRepository[T any] interface {
	Save(ctx Context[T], side *entity.Side) (*entity.Side, error)
	FindByID(ctx Context[T], id uuid.UUID) (*entity.Side, error)
}

type GormSideRepository struct {}

func (g *GormSideRepository) Save(ctx Context[*gorm.DB], side *entity.Side) (*entity.Side, error) {
	if err :=  ctx.Executor().Save(side).Error; err != nil {
		return nil, err
	}
	return side, nil
}

func (g *GormSideRepository) FindByID(ctx Context[*gorm.DB], id uuid.UUID) (*entity.Side, error) {
	side := new(entity.Side)

	err := ctx.Executor().First(side, "id = ?", id).Error;

	if err == nil {
		return side, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return nil, err
}

func NewGormSideRepository() *GormSideRepository {
	return &GormSideRepository{}
}