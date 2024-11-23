package repository

import (
	"context"

	"gorm.io/gorm"
)

type Context[T any] interface {
	context.Context
	Executor() T
}

type TxContext[T any] interface {
	Context[T]
	Commit() error
	Rollback() error
}

type GormContext struct {
	context.Context
	executor *gorm.DB
}

func (c *GormContext) Executor() *gorm.DB {
	return c.executor
}

type GormTxContext struct {
	context.Context
	executor *gorm.DB
}

func (tx *GormTxContext) Executor() *gorm.DB {
	return tx.executor
}

func (tx *GormTxContext) Commit() error {
	return tx.executor.Commit().Error
}

func (tx *GormTxContext) Rollback() error {
	return tx.executor.Rollback().Error
}

