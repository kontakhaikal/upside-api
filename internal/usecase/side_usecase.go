package usecase

import (
	"context"

	"github.com/fkrhykal/upside-api/internal/dto"
	"github.com/fkrhykal/upside-api/internal/entity"
	"github.com/fkrhykal/upside-api/internal/errors"
	"github.com/fkrhykal/upside-api/internal/repository"
	"github.com/fkrhykal/upside-api/internal/util"
	"github.com/google/uuid"
	"gorm.io/gorm"
)


type SideUseCase interface {
	Create(ctx context.Context, req *dto.CreateSideRequest) (*dto.CreateSideResponse, error)
	Join(ctx context.Context, req *dto.JoinSideRequest) (*dto.JoinSideResponse, error)
}

type SideUseCaseImpl struct {
	validator util.Validator
	sideRepository repository.SideRepository[*gorm.DB]
	membershipRepository repository.MembershipRepository[*gorm.DB]
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

	membership := &entity.Membership{
		ID: uuid.New(),
		UserID: req.UserID,
		SideID: side.ID,
		Role: entity.Author,
	}

	membership, err = s.membershipRepository.Save(repoCtx, membership)

	if err != nil {
		return nil, err
	}

	if err = repoCtx.Commit(); err != nil {
		return nil, err
	}

	return &dto.CreateSideResponse{
		SideID: side.ID,
	}, nil
}

func (s *SideUseCaseImpl) Join(ctx context.Context, req *dto.JoinSideRequest) (*dto.JoinSideResponse, error) {
	if err := s.validator.ValidateDTO(req); err != nil {
		return nil, err
	}

	repoCtx := s.contextManager.WithoutTx(ctx)

	side, err := s.sideRepository.FindByID(repoCtx, req.SideID)
	
	if err != nil {
		return nil, err
	}

	if side == nil {
		return nil, errors.ErrSideNotFound
	}

	membership, err := s.membershipRepository.FindBySideIDAndUserID(repoCtx, side.ID, req.UserID)

	if err != nil {
		return nil, err
	}

	if membership != nil {
		return nil, errors.ErrAlreadyJoinedSide
	}

	membership = &entity.Membership{
		ID: uuid.New(),
		UserID: req.UserID,
		SideID: side.ID,
		Role: entity.Member,
	}

	membership, err = s.membershipRepository.Save(repoCtx, membership)

	if err != nil {
		return nil, err
	}

	return &dto.JoinSideResponse{MembershipID: membership.ID}, nil
}

func NewSideUseCase(
	validator util.Validator,
	sideRepository repository.SideRepository[*gorm.DB],
	membershipRepository repository.MembershipRepository[*gorm.DB],
	contextManager repository.ContextManager[*gorm.DB],
) SideUseCase {
	return &SideUseCaseImpl{
		validator,
		sideRepository,
		membershipRepository,
		contextManager,
	}
} 