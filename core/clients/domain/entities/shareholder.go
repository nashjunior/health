package entities

import (
	"health/core/application/entities"
	valueobjects "health/core/application/value-objects"
)

type Shareholder struct {
	entities.Entity

	owner      *Person
	enterprise *Enterprise
}

func (shareholder *Shareholder) GetEnterprise() *Enterprise {
	return shareholder.enterprise
}

func (shareholder *Shareholder) GetOwner() *Person {
	return shareholder.owner
}

func NewShareHolder(owner *Person, enterprise *Enterprise, id *valueobjects.UniqueEntityUUID) (*Shareholder, error) {
	return &Shareholder{
		Entity:     entities.NewEntity(id, nil),
		owner:      owner,
		enterprise: enterprise,
	}, nil

}
