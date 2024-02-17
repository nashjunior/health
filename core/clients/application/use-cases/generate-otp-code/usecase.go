package generateotpcode

import (
	"errors"
	applicationErrors "health/core/application/errors"
	"health/core/clients/domain/entities"
	"health/core/clients/domain/repositories"
	"health/core/infra/utils"
	"time"
)

type Input struct {
	IdUser string
}

type Usecase struct {
	usersRepository            repositories.IUsersRepository
	personsRepository          repositories.IPersonsRepository
	validationsCodesRepository repositories.IValidationsCodesRepository
}

func (usecase *Usecase) Execute(Input Input) (*string, error) {
	user, err := usecase.usersRepository.FindByID(Input.IdUser, nil)

	if err != nil {
		return nil, err
	}

	validationCode, err := usecase.validationsCodesRepository.FindCurrentUserValidationCode(*user)

	switch err.(type) {
	case *applicationErrors.NotFoundError:
		{
			code := utils.GenerateCode(6)
			now := time.Now().Format(time.RFC3339)

			validationCode, err = entities.NewValidationCode(&code, &now, user, nil)

			if err != nil {
				return nil, err
			}

			err = usecase.validationsCodesRepository.Create(*validationCode, nil)

			if err != nil {
				return nil, err
			}

			return &code, nil
		}

	default:
		{
			return nil, errors.New("Theres a code still active")
		}
	}

}
