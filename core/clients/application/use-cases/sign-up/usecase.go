package signup

import (
	"fmt"
	"health/core/clients/domain/entities"
	"health/core/clients/domain/repositories"
	"time"
)

type Input struct {
	Name         string
	BirthdayDate *time.Time
}

type Output struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type Usecase struct {
	usersRepository repositories.IUsersRepository
}

func (usecase *Usecase) Execute(input Input) (*Output, error) {
	fmt.Println("Create user usecase")
	user, err := entities.NewUser(&input.Name, nil)

	if err != nil {
		return nil, err
	}

	usecase.usersRepository.Create(*user, nil)

	fmt.Println("Should create a otp here")

	return &Output{
		Id:        user.GetID().String(),
		Name:      user.GetName(),
		CreatedAt: user.CreatedAt,
	}, nil
}

func NewUsecase(usersRepository repositories.IUsersRepository) Usecase {
	return Usecase{
		usersRepository: usersRepository,
	}
}
