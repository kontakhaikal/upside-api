package repository

import (
	"github.com/fkrhykal/upside-api/internal/entity"
	"gorm.io/gorm"
)

type PostRepository[T any] interface {
	Save(ctx Context[T], post *entity.Post) (*entity.Post, error)
}

type GormPostRepository struct {}


func (g *GormPostRepository) Save(ctx Context[*gorm.DB], post *entity.Post) (*entity.Post, error) {
	if err := ctx.Executor().Save(post).Error; err != nil {
		return nil, err
	}
	return post, nil
}

func NewGormPostRepository() *GormPostRepository {
	return &GormPostRepository{}
}