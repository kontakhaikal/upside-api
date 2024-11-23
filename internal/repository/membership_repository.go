package repository

import (
	"errors"

	"github.com/fkrhykal/upside-api/internal/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MembershipRepository[T any] interface {
	Save(ctx Context[T], membership *entity.Membership) (*entity.Membership, error)
	FindBySideIDAndUserID(ctx Context[T], sideID uuid.UUID, userID uuid.UUID) (*entity.Membership, error)
}

type GormMembershipRepository struct{}

func (g *GormMembershipRepository) Save(ctx Context[*gorm.DB], membership *entity.Membership) (*entity.Membership, error) {
	if err := ctx.Executor().Save(membership).Error; err != nil {
		return nil, err
	}
	return membership, nil
}

func (g *GormMembershipRepository) FindBySideIDAndUserID(ctx Context[*gorm.DB], sideID uuid.UUID, userID uuid.UUID) (*entity.Membership, error) {
	membership := new(entity.Membership)

	err := ctx.Executor().Where("side_id = ?", sideID).Where("user_id = ?", userID).First(membership).Error

	if err == nil {
		return membership, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return nil, err
}

func NewGormMembershipRepository() *GormMembershipRepository {
	return &GormMembershipRepository{}
}
