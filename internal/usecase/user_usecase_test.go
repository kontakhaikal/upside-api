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
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


type userUseCaseDep struct {
	repository repository.UserRepository[*gorm.DB]
	useCase usecase.UserUseCase
	db *gorm.DB
	ctxManager repository.ContextManager[*gorm.DB]
}

func setupUserUseCase(t *testing.T) *userUseCaseDep {
	db, err := gorm.Open(sqlite.Open("file:test?mode=memory&cache=shared"))

	if err != nil {
		t.Fatalf("database connection error: %+v",err)
	}

	if err = db.AutoMigrate(&entity.User{}); err != nil {
		t.Fatalf("automigration error: %+v", err)
	}

	validator := config.NewPlaygoundValidator()
	contextManager := repository.NewGormContextManager(db)
	userRepository := repository.NewGormUserRepository()
	passwordHasher := config.NewBcryptPasswordHasher()
	credentialUtil := config.NewJwtCredentialUtil([]byte("secret"))
	useCase := usecase.NewUserService(
		validator,
		userRepository,
		contextManager,
		passwordHasher,
		credentialUtil,
	)

	return &userUseCaseDep{
		repository: userRepository,
		useCase: useCase,
		db: db,
		ctxManager: contextManager,
	}
}


func TestRegisterUserUseCase(t *testing.T) {
	
	dep := setupUserUseCase(t)

	ctx := context.TODO()

	req :=  &dto.RegisterUserRequest{
		FirstName: faker.FirstName(),
		LastName: faker.LastName(),
		Username: faker.Username(),
		Password: faker.Password(),
	}

	res, err := dep.useCase.Register(ctx, req)

	if err != nil {
		t.Fatalf("register user error: %+v", err)
	}

	user := entity.User{}

	if err = dep.db.Model(&user).First(&user, "id = ?", res.ID).Error; err != nil {
		t.Fatalf("error when check database: %+v", err)
	}

	if res.ID != user.ID.String() {
		t.Fatalf("mismatch user id %s with result id %s", user.ID, res.ID)
	}
}

func TestLoginUserUseCase(t *testing.T) {
	dep := setupUserUseCase(t)

	ctx := context.TODO()

	req :=  &dto.RegisterUserRequest{
		FirstName: faker.FirstName(),
		LastName: faker.LastName(),
		Username: faker.Username(),
		Password: faker.Password(),
	}

	_, err := dep.useCase.Register(ctx, req)

	if err != nil {
		t.Fatalf("register user error: %+v", err)
	}

	res, err := dep.useCase.Login(ctx, &dto.LoginUserRequest{Username: req.Username, Password: req.Password})

	if err != nil {
		t.Fatalf("user login error: %+v", err)
	}

	t.Logf("%+v", res)
}