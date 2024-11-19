package repository

import (
	"context"

	"gorm.io/gorm"
)

type ContextManager[T any] interface {
	// create context with already begin transaction
	WithTx(context.Context) TxContext[T]
	WithoutTx(context.Context) Context[T]
}


type GormContextManager struct {
	db *gorm.DB
}

func (cm *GormContextManager) WithTx(parent context.Context) TxContext[*gorm.DB]{
	tx := cm.db.Begin()
	return &GormTxContext{
		Context: parent,
		executor: tx,
	}
}

func (cm *GormContextManager) WithoutTx(parent context.Context) Context[*gorm.DB] {
	return &GormContext{
		Context: parent,
		executor: cm.db,
	}
}

func NewGormContextManager(db *gorm.DB) ContextManager[*gorm.DB] {
	return &GormContextManager{db}
}