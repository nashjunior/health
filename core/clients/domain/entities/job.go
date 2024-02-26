package entities

import (
	"health/core/application/entities"
	valueobjects "health/core/application/value-objects"

	"github.com/go-playground/validator/v10"
)

var validationJob *validator.Validate

type Job struct {
	entities.Entity
	id *int

	name *string

	managers     *[]Job
	subordinates *[]Job
}

func (job *Job) SetManagers(managers []Job) {
	job.managers = &managers
}

func (job *Job) GetManagers() []Job {
	return *job.managers
}

func (job *Job) SetSubordinates(subordinates []Job) {
	job.subordinates = &subordinates
}

func (job *Job) GetSubordinates() []Job {
	return *job.subordinates
}

func (job *Job) SetInternalId(id int) {
	job.id = &id
}

func (job *Job) GetInternalId() int {
	return *job.id
}

func (job *Job) setName(name string) error {
	err := validationJob.Var(name, "required")

	if err != nil {
		return err
	}

	job.name = &name
	return nil
}

func (job *Job) GetName() string {
	return *job.name
}

func (job *Job) Update(name *string) error {
	if name != nil {
		err := job.setName(*name)
		if err != nil {
			return err
		}
	}

	return nil
}

func NewJob(
	name *string,
	managers *[]Job,
	subordinates *[]Job,
	id *valueobjects.UniqueEntityUUID,
) (*Job, error) {
	job := &Job{}
	validationJob = validator.New()

	if name != nil {
		err := job.setName(*name)
		if err != nil {
			return nil, err
		}
	}

	job.managers = managers
	job.subordinates = subordinates

	job.Entity = entities.NewEntity(id, nil)
	return job, nil
}
