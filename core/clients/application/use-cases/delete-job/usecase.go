package deletejob

import (
	"health/core/clients/domain/repositories"
)

type Input struct {
	Id string
}

type Usecase struct {
	jobsRepository            repositories.IJobsRepository
	jobsHierarchiesRepository repositories.IJobsHierarchiesRepository
}

func (usecase *Usecase) Execute(input Input) error {

	job, err := usecase.jobsRepository.FindByID(input.Id, nil)

	if err != nil {
		return err
	}

	usecase.jobsHierarchiesRepository.DeleteByJob(*job, nil)
	usecase.jobsRepository.Delete(*job, nil)

	return nil
}

func New(
	jobsRepository repositories.IJobsRepository,
	jobsHierarchiesRepository repositories.IJobsHierarchiesRepository,
) Usecase {
	return Usecase{
		jobsRepository:            jobsRepository,
		jobsHierarchiesRepository: jobsHierarchiesRepository,
	}
}
