package usecase

import (
	"context"

	"github.com/fkrhykal/upside-api/internal/dto"
	"github.com/fkrhykal/upside-api/internal/entity"
	"github.com/fkrhykal/upside-api/internal/repository"
	"github.com/fkrhykal/upside-api/internal/util"
	"github.com/google/uuid"
	"gorm.io/gorm"
)


type SideUseCase interface {
	Create(ctx context.Context, req *dto.CreateSideRequest) (*dto.CreateSideResponse, error)
}

type SideUseCaseImpl struct {
	validator util.Validator
	sideRepository repository.SideRepository[*gorm.DB]
	membershipRepository repository.MemberShipRepository[*gorm.DB]
	contextManager repository.ContextManager[*gorm.DB]
}

func (s *SideUseCaseImpl) Create(ctx context.Context, req *dto.CreateSideRequest) (*dto.CreateSideResponse, error) {
	if err := s.validator.ValidateDTO(req); err != nil {
		return nil, err
	}

	side := &entity.Side{
		ID: uuid.New(),
		Name: req.Name,
	}

	repoCtx := s.contextManager.WithTx(ctx)

	defer repoCtx.Rollback()

	side, err := s.sideRepository.Save(repoCtx, side)

	if err != nil {
		return nil, err
	}

	membership := &entity.MemberShip{
		ID: uuid.New(),
		UserID: req.UserID,
		Side: side.ID,
		Role: entity.Author,
	}

	membership, err = s.membershipRepository.Save(repoCtx, membership)

	if err != nil {
		return nil, err
	}

	repoCtx.Commit()

	return &dto.CreateSideResponse{
		ID: side.ID,
	}, nil
}

func NewSideUseCase(
	validator util.Validator,
	sideRepository repository.SideRepository[*gorm.DB],
	membershipRepository repository.MemberShipRepository[*gorm.DB],
	contextManager repository.ContextManager[*gorm.DB],
) SideUseCase {
	return &SideUseCaseImpl{
		validator,
		sideRepository,
		membershipRepository,
		contextManager,
	}
} 