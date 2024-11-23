package usecase_test

import (
	"context"
	"fmt"
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
	"gorm.io/gorm/logger"
)

type sideUseCaseDep struct {
	sideRepo      repository.SideRepository[*gorm.DB]
	membershipRep repository.MembershipRepository[*gorm.DB]
	ctxManager    repository.ContextManager[*gorm.DB]
	useCase       usecase.SideUseCase
}

func setupSideUseCase(t *testing.T) *sideUseCaseDep {
	db, err := gorm.Open(sqlite.Open(fmt.Sprintf("file:../../db/%s.db", uuid.New())), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		t.Fatalf("database connection error: %+v", err)
	}

	if err = db.AutoMigrate(&entity.Side{}, &entity.Membership{}); err != nil {
		t.Fatalf("automigration error: %+v", err)
	}

	ctxManager := repository.NewGormContextManager(db)
	sideRepo := repository.NewGormSideRepository()
	membershipRepo := repository.NewGormMembershipRepository()
	validator := config.NewPlaygoundValidator()
	useCase := usecase.NewSideUseCase(
		validator,
		sideRepo,
		membershipRepo,
		ctxManager,
	)

	return &sideUseCaseDep{
		sideRepo,
		membershipRepo,
		ctxManager,
		useCase,
	}
}

func TestCreateSide(t *testing.T) {
	dep := setupSideUseCase(t)

	ctx := context.TODO()

	req := &dto.CreateSideRequest{UserID: uuid.New(), Name: faker.Name()}

	res, err := dep.useCase.Create(ctx, req)

	if err != nil {
		t.Fatalf("error on creating side: %+v", err)
	}

	t.Logf("res id: %s", res.SideID)
}

func TestJoinSide(t *testing.T) {
	dep := setupSideUseCase(t)

	ctx := context.TODO()

	req := &dto.CreateSideRequest{UserID: uuid.New(), Name: faker.Name()}

	r, err := dep.useCase.Create(ctx, req)

	if err != nil {
		t.Fatalf("error when creating side: %+v", err)
	}

	res, err := dep.useCase.Join(ctx, &dto.JoinSideRequest{SideID: r.SideID, UserID: uuid.New()})

	if err != nil {
		t.Fatalf("error on join side: %+v", err)
	}

	t.Logf("membership id: %s", res.MembershipID)
}
