package entities

import (
	"health/core/application/entities"
	valueobjects "health/core/application/value-objects"
	"time"

	"github.com/go-playground/validator/v10"
)

var validationEnterprise *validator.Validate

type Enterprise struct {
	entities.Entity

	cnpj         *string
	socialReason *string

	user        *User
	shareolders *[]Shareholder
}

func (enterprise *Enterprise) setCNPJ(cnpj string) error {
	err := validationEnterprise.Var(cnpj, "required,len=14")

	if err != nil {
		return err
	}

	enterprise.cnpj = &cnpj

	return nil
}

func (enterprise *Enterprise) GetCnpj() *string {
	return enterprise.cnpj
}

func (enterprise *Enterprise) GetSocialReason() *string {
	return enterprise.socialReason
}

func (enterprise *Enterprise) setSocialReason(socialReason string) error {
	err := validationEnterprise.Var(socialReason, "required,min=3")

	if err != nil {
		return err
	}

	enterprise.socialReason = &socialReason

	return nil
}

func (enterprise *Enterprise) GetUser() User {
	return *enterprise.user
}

func (enterprise *Enterprise) Update(cnpj *string, socialReason *string) error {
	var err error

	if cnpj != nil {
		err = enterprise.setCNPJ(*cnpj)
		if err != nil {
			return err
		}

	}

	if socialReason != nil {
		err = enterprise.setSocialReason(*socialReason)
		if err != nil {
			return err
		}
	}

	now := time.Now()
	enterprise.UpdatedAt = &now
	return nil

}

func NewEnterprise(cnpj *string, socialReason *string, user *User, id *valueobjects.UniqueEntityUUID) (*Enterprise, error) {
	validationEnterprise = validator.New(validator.WithRequiredStructEnabled())
	var err error

	enterprise := &Enterprise{}

	if cnpj != nil {
		err = enterprise.setCNPJ(*cnpj)
		if err != nil {
			return nil, err
		}

	}

	if socialReason != nil {
		err = enterprise.setSocialReason(*socialReason)
		if err != nil {
			return nil, err
		}
	}

	enterprise.Entity = entities.NewEntity(id, nil)
	enterprise.user = user

	return enterprise, nil
}
