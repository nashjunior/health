package repositories

import (
	"health/core/application/entities"
	"math/big"

	"github.com/google/uuid"
)

type IResponseSearchable[T entities.Entity] struct {
	Total big.Int
	Items T
}

type SearchablePagination struct {
	Page    int
	PerPage int
}

type SearchableQuery struct {
	Value  string
	Fields []string
}

type SearchableRepository[T entities.Entity, Filter any] interface {
	FindByUUID(id uuid.UUID) (T, error)
	FindByID(id string) (T, error)
	Find() []T
	FindAndCount() IResponseSearchable[T]
}

type InsertableRepository[T entities.Entity] interface {
	Create(entitiy T) error
	CreateMany(entitiy []T) error
}

type UpdatableRepository[T entities.Entity] interface {
	Update(entity T) error
}

type DeletableRepository[T entities.Entity] interface {
	Delete(entity T) error
	DeleteMany(entidades []T) error
}
