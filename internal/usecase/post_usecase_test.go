package usecase_test

import (
	"context"
	"testing"

	"github.com/fkrhykal/upside-api/internal/config"
	"github.com/fkrhykal/upside-api/internal/dto"
	"github.com/fkrhykal/upside-api/internal/entity"
	"github.com/fkrhykal/upside-api/internal/repository"
	"github.com/fkrhykal/upside-api/internal/usecase"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type postUseCaseDep struct {
	sideRepo repository.SideRepository[*gorm.DB]
	membershipRep repository.MembershipRepository[*gorm.DB]
	postRep repository.PostRepository[*gorm.DB]
	ctxManager repository.ContextManager[*gorm.DB]
	useCase usecase.PostUseCase
	db *gorm.DB
}

func setupPostUseCase(t *testing.T) *postUseCaseDep {
	db, err := gorm.Open(sqlite.Open("file:test?mode=memory&cached=shared"))

	if err != nil {
		t.Fatalf("database connection error: %+v",err)
	}

	if err = db.AutoMigrate(&entity.Side{}, &entity.Membership{}, &entity.Post{}); err != nil {
		t.Fatalf("automigration error: %+v", err)
	}

	ctxManager := repository.NewGormContextManager(db)
	sideRepo := repository.NewGormSideRepository()
	membershipRepo := repository.NewGormMembershipRepository()
	postRepo := repository.NewGormPostRepository()
	validator := config.NewPlaygoundValidator()
	useCase := usecase.NewPostUseCase(
		validator,
		postRepo,
		membershipRepo,
		sideRepo,
		ctxManager,
	)

	return &postUseCaseDep{
		sideRepo,
		membershipRepo,
		postRepo,
		ctxManager,
		useCase,
		db,
	}
}

func TestCreatePost(t *testing.T) {
	dep := setupPostUseCase(t)

	ctx := context.TODO()

	req := &dto.CreatePostRequest{
		Body: faker.Sentence(),
		SideID: uuid.New(),
		AuthorID: uuid.New(),
	}

	if err := dep.db.Save(&entity.Side{
		ID: req.SideID,
		Name: faker.Name(),
	}).Error; err != nil {
		t.Fatalf("error when create side %+v", err)
	}

	if err := dep.db.Save(&entity.Membership{
		ID: uuid.New(),
		UserID: req.AuthorID,
		SideID: req.SideID,
		Role: entity.Member,
	}).Error; err != nil {
		t.Fatalf("error when create membership %+v", err)
	}

	res, err := dep.useCase.Create(ctx, req)

	if err != nil {
		t.Fatalf("error when create post: %+v", err)
	}

	t.Logf("created post id: %s", res.PostID)
}