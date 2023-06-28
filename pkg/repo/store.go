package repo

import (
	"gorm.io/gorm"
)

// FinallyFunc function to finish a transaction
type FinallyFunc = func(error) error

// Store persistent data interface
type Store interface {
	DB() *gorm.DB
	NewTransaction() (repo *Repo, finalFn IFinalFunc)
	Shutdown() error
}

type IFinalFunc interface {
	Commit() error
	Rollback(err error) error
}
