package entities

import (
	"health/core/application/entities"
	valueobjects "health/core/application/value-objects"

	"github.com/go-playground/validator/v10"
)

var validationJobHierarchy *validator.Validate

type JobHierarchy struct {
	entities.Entity

	id *int

	job     *Job
	manager *Job
}

func (job *JobHierarchy) GetJob() Job {
	return *job.job
}

func (job *JobHierarchy) GetManager() *Job {
	return job.manager
}

func (job *JobHierarchy) Update(manager *Job) {
	if manager != nil {
		job.manager = manager
	}
}

func NewJobHierarchy(
	job *Job,
	manager *Job,
	id *valueobjects.UniqueEntityUUID,
) (*JobHierarchy, error) {
	jobTree := &JobHierarchy{Entity: entities.NewEntity(id, nil)}

	jobTree.job = job
	jobTree.manager = manager
	return jobTree, nil
}
