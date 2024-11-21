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

type UserUseCase interface {
	Register(ctx context.Context, req *dto.RegisterUserRequest) (*dto.RegisterUserResponse, error)
	Login(ctx context.Context, req *dto.LoginUserRequest) (*dto.LoginUserResponse, error)
}

type UserUseCaseImpl struct {
	validator util.Validator
	userRepository repository.UserRepository[*gorm.DB]
	contextManager repository.ContextManager[*gorm.DB]
	passwordHasher util.PasswordHasher
	credetialUtil util.CredentialUtil
}

func (u *UserUseCaseImpl) Register(ctx context.Context, req *dto.RegisterUserRequest) (*dto.RegisterUserResponse, error) {
	if err:= u.validator.ValidateDTO(req); err != nil {
		return nil, err
	}

	repoCtx := u.contextManager.WithoutTx(ctx)

	exist, err := u.userRepository.UsernameExist(repoCtx, req.Username); 

	if err != nil {
		return nil, err
	}
	
	if exist {
		return nil, errors.ErrUsernameUsed
	}

	hashedPassword, err := u.passwordHasher.Hash(req.Password)
	
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		ID: uuid.New(),
		FirstName: req.FirstName,
		LastName: req.LastName,
		Username: req.Username,
		Password: hashedPassword,
	}

	user, err = u.userRepository.Save(repoCtx, user)

	if err != nil {
		return nil, err
	}

	return &dto.RegisterUserResponse{ID: user.ID.String()}, nil
}

func (u *UserUseCaseImpl) Login(ctx context.Context, req *dto.LoginUserRequest) (*dto.LoginUserResponse, error) {
	if err := u.validator.ValidateDTO(req); err != nil {
		return nil, errors.ErrAuthentication
	}

	repoCtx := u.contextManager.WithoutTx(ctx)

	user, err := u.userRepository.FindByUsername(repoCtx, req.Username)

	if err != nil {
		return nil, err
	}

	credential, err := u.credetialUtil.GenerateToken(
		&dto.UserCredential{
			ID: user.ID,
		},
	)

	if err != nil {
		return nil, err
	}

	return &dto.LoginUserResponse{CredentialToken: credential}, nil
}

func NewUserService(
	validator util.Validator,
	userRepository repository.UserRepository[*gorm.DB],
	contextManager repository.ContextManager[*gorm.DB],
	passwordHasher util.PasswordHasher,
	credentialUtil util.CredentialUtil,
	) *UserUseCaseImpl {
	return &UserUseCaseImpl{
		validator,
		userRepository,
		contextManager,
		passwordHasher,
		credentialUtil,
	}
}