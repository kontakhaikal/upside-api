package repository

import (
	"errors"

	"github.com/fkrhykal/upside-api/internal/entity"
	"gorm.io/gorm"
)


type UserRepository[T any] interface {
	UsernameExist(ctx Context[T], username string) (bool, error)
	FindByUsername(ctx Context[T], username string) (*entity.User, error)
	Save(ctx Context[T], user *entity.User) (*entity.User, error)
}


type GormUserRepository struct {}

func (g *GormUserRepository) UsernameExist(ctx Context[*gorm.DB], username string) (bool, error) {
	var count int64


	if err := ctx.Executor().Model(&entity.User{}).Where("username = ?", username).Count(&count).Error; err != nil {
		return false, err
	}



	return count != 0, nil
}

func (g *GormUserRepository) Save(ctx Context[*gorm.DB], user *entity.User) (*entity.User, error) {
	if err := ctx.Executor().Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (g *GormUserRepository) FindByUsername(ctx Context[*gorm.DB], username string) (*entity.User, error) {
	user := new(entity.User)

	err := ctx.Executor().First(&user, "username = ?", username).Error

	if err == nil {
		return user, nil
	}
	
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return nil, err
}

func NewGormUserRepository() UserRepository[*gorm.DB] {
	return &GormUserRepository{}
}