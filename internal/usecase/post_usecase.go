package usecase

import (
	"context"
	"time"

	"github.com/fkrhykal/upside-api/internal/dto"
	"github.com/fkrhykal/upside-api/internal/entity"
	"github.com/fkrhykal/upside-api/internal/errors"
	"github.com/fkrhykal/upside-api/internal/repository"
	"github.com/fkrhykal/upside-api/internal/util"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostUseCase interface {
	Create(ctx context.Context, req *dto.CreatePostRequest) (*dto.CreatePostResponse, error)
}

type PostUseCaseImpl struct {
	validator util.Validator
	postRepository repository.PostRepository[*gorm.DB]
	memberRepository repository.MembershipRepository[*gorm.DB]
	sideRepository repository.SideRepository[*gorm.DB]
	contextManager repository.ContextManager[*gorm.DB]
}

func (p *PostUseCaseImpl) Create(ctx context.Context, req *dto.CreatePostRequest) (*dto.CreatePostResponse, error) {
	if err := p.validator.ValidateDTO(req); err != nil {
		return nil, err
	}

	repoCtx := p.contextManager.WithoutTx(ctx)

	side, err := p.sideRepository.FindByID(repoCtx, req.SideID)

	if err != nil {
		return nil, err
	}

	if side == nil {
		return nil, errors.ErrSideNotFound
	}

	membership, err  := p.memberRepository.FindBySideIDAndUserID(repoCtx, side.ID, req.AuthorID)

	if err != nil {
		return nil, err
	}

	if membership == nil {
		return nil, errors.ErrNotAMember
	}

	post := &entity.Post{
		ID: uuid.New(),
		Body: req.Body,
		CreatedAt: time.Now(),
		LastUpdated: time.Now(),
		AuthorID: req.AuthorID,
		SideID: side.ID,
	}

	post, err = p.postRepository.Save(repoCtx, post)

	if err != nil {
		return nil, err
	}

	return &dto.CreatePostResponse{PostID: post.ID}, nil
}

func NewPostUseCase(
	validator util.Validator,
	postRepository repository.PostRepository[*gorm.DB],
	memberRepository repository.MembershipRepository[*gorm.DB],
	sideRepository repository.SideRepository[*gorm.DB],
	contextManager repository.ContextManager[*gorm.DB],

) PostUseCase {
	return &PostUseCaseImpl{
		validator,
		postRepository,
		memberRepository,
		sideRepository,
		contextManager,
	}
}