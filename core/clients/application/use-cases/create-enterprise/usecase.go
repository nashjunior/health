package createenterprise

import (
	"health/core/clients/application/services"
	"time"
)

type Input struct {
	Name         string
	Cnpj         string
	Email        *string
	SocialReason string
	BirthdayDate *time.Time
}

type Output struct {
	Id           *string    `json:"id"`
	Name         string     `json:"name"`
	Cnpj         string     `json:"cnpj"`
	SocialReason string     `json:"social_reason"`
	BirthdayDate *time.Time `json:"birthday_date"`
	CreatedAt    time.Time  `json:"created_at"`
}

type Usecase struct{}

func (usecase *Usecase) Execute(input Input) (*Output, error) {

	var output Output

	var formatedBirthday *string

	if input.BirthdayDate != nil {
		formated := input.BirthdayDate.Format(time.RFC3339)
		formatedBirthday = &formated
	}

	err := services.CreateKeycloakUser(services.KeycloakUser{
		Username:      input.Name,
		Email:         *input.Email,
		EmailVerified: true,
		FirstName:     input.Name,
		LastName:      input.Name,
		Enabled:       true,
		Attributes: map[string]*string{
			"cnpj":          &input.Cnpj,
			"socialReason":  &input.SocialReason,
			"birthday_date": formatedBirthday,
		},
	})

	if err != nil {
		return nil, err
	}

	output.Name = input.Name
	output.Cnpj = input.Cnpj
	output.SocialReason = input.SocialReason
	output.CreatedAt = time.Now()

	return &output, nil
}

func New() Usecase {
	return Usecase{}
}
