package tests

import (
	"health/core/application/entities"
	"health/core/application/repositories"
)

type StubEntity struct {
	entities.Entity
	Name string
	Age  uint8
}

type StubInMemoryRepository struct {
	repositories.AbstractIntMemoryRepository
}

func Create() repositories.InsertableRepository[StubEntity] {
	return repositories.AbstractIntMemoryRepository{}
}

func A() {
	memory = StubInMemoryRepository{AbstractIntMemoryRepository: Create()}
}
