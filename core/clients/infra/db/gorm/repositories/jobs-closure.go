package repositories

import (
	"health/core/application/errors"
	repo "health/core/application/repositories"
	ent "health/core/clients/domain/entities"
	"health/core/clients/domain/repositories"
	"health/core/clients/infra/db/gorm/entities"
	"health/core/infra/db/gorm"
	"math/big"

	valueobjects "health/core/application/value-objects"

	"github.com/google/uuid"
	orm "gorm.io/gorm"
)

type JobsHierarchiesGorm struct {
	db          *orm.DB
	jobClosure  *entities.JobHierarchy
	jobsClosure *[]entities.JobHierarchy
}

func (repo *JobsHierarchiesGorm) FindByUUID(id uuid.UUID, conn repo.IConnection) (*ent.JobHierarchy, error) {
	if conn2, ok := conn.(*gorm.ConnectionGorm); ok && conn2 != nil {
		repo.db = conn2.Tx
	} else {
		repo.db = gorm.Connection.Db
	}

	if result := repo.db.Omit("Id").InnerJoins("Job").First(&repo.jobClosure, "jobs_closure.uuid = ?", id); result.Error != nil {
		return nil, result.Error
	}

	if repo.jobClosure == nil {
		return nil, errors.NewNotFoundError("Could not found error using id")
	}

	jobGorm := repo.jobClosure.Job
	uniqueIdJob := valueobjects.NewUniqueUUID(uuid.NullUUID{UUID: jobGorm.Uuid})

	job, _ := ent.NewJob(&jobGorm.Name, nil, nil, &uniqueIdJob)

	uniqueTypeTransactionId := valueobjects.NewUniqueUUID(uuid.NullUUID{UUID: repo.jobClosure.Uuid})

	return ent.NewJobHierarchy(job, nil, &uniqueTypeTransactionId)
}

func (repo *JobsHierarchiesGorm) FindByID(id string, conn repo.IConnection) (*ent.JobHierarchy, error) {
	parsedId, err := uuid.Parse(id)

	if err != nil {
		return nil, err
	}

	return repo.FindByUUID(parsedId, conn)
}

func (repo *JobsHierarchiesGorm) Find(params *repositories.SearchParamJobHierarchy, conn repo.IConnection) []ent.JobHierarchy {
	if conn2, ok := conn.(*gorm.ConnectionGorm); ok && conn2 != nil {
		repo.db = conn2.Tx
	} else {
		repo.db = gorm.Connection.Db
	}

	if result := repo.db.Omit("id").InnerJoins("Job").Find(&repo.jobsClosure); result.Error != nil {
		return []ent.JobHierarchy{}
	}

	var entities []ent.JobHierarchy

	for _, item := range *repo.jobsClosure {
		job := item.Job

		uniqueJobId := valueobjects.NewUniqueUUID(uuid.NullUUID{UUID: job.Uuid})
		jobEntity, _ := ent.NewJob(&job.Name, nil, nil, &uniqueJobId)

		uniqueId := valueobjects.NewUniqueUUID(uuid.NullUUID{UUID: item.Uuid})

		ent, err := ent.NewJobHierarchy(
			jobEntity,
			nil,
			&uniqueId,
		)

		if err == nil {
			entities = append(entities, *ent)
		}
	}

	return entities
}

func (repo *JobsHierarchiesGorm) FindAndCount(params *repositories.SearchParamJobHierarchy, conn repo.IConnection) repositories.IResponseSearchableJobsHirearchy {
	items := repo.Find(params, conn)

	return repositories.IResponseSearchableJobsHirearchy{
		Total: *big.NewInt(int64(len(items))),
		Items: items,
	}
}

func (repo *JobsHierarchiesGorm) Create(entity *ent.JobHierarchy, conn repo.IConnection) error {

	job := entity.GetJob()
	manager := entity.GetManager()

	var idManager *int

	if manager != nil {
		internalId := manager.GetInternalId()
		idManager = &internalId
	}

	ent := entities.JobHierarchy{
		Uuid:      entity.GetID(),
		IdJob:     job.GetInternalId(),
		IdManager: idManager,
		CreatedAt: entity.CreatedAt,
	}

	ommitFields := []string{"Id", "DeletedAt", "UpdatedAt"}

	if conn2, ok := conn.(*gorm.ConnectionGorm); ok && conn2 != nil {
		repo.db = conn2.Tx
	} else {
		repo.db = gorm.Connection.Db
	}
	if result := repo.db.Omit(ommitFields...).Create(&ent); result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *JobsHierarchiesGorm) CreateMany(newEntities []ent.JobHierarchy, conn repo.IConnection) error {

	var hierarchies []entities.JobHierarchy

	for _, item := range newEntities {
		job := item.GetJob()
		manager := item.GetManager()
		var idManager *int

		if manager != nil {
			id := manager.GetInternalId()
			idManager = &id
		}

		hierarchies = append(hierarchies, entities.JobHierarchy{
			Uuid:      item.GetID(),
			IdJob:     job.GetInternalId(),
			IdManager: idManager,
			CreatedAt: item.CreatedAt,
		})
	}

	ommitFields := []string{"Id", "DeletedAt", "UpdatedAt"}

	if conn2, ok := conn.(*gorm.ConnectionGorm); ok && conn2 != nil {
		repo.db = conn2.Tx
	} else {
		repo.db = gorm.Connection.Db
	}
	if result := repo.db.Omit(ommitFields...).Create(&hierarchies); result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *JobsHierarchiesGorm) Update(entity ent.JobHierarchy, conn repo.IConnection) error {

	job := entity.GetJob()
	manager := entity.GetManager()

	var idManager *int

	if manager != nil {
		internalId := manager.GetInternalId()
		idManager = &internalId
	}

	ent := entities.JobHierarchy{
		Uuid:      entity.GetID(),
		IdJob:     job.GetInternalId(),
		IdManager: idManager,
		CreatedAt: entity.CreatedAt,
	}
	if conn2, ok := conn.(*gorm.ConnectionGorm); ok && conn2 != nil {
		repo.db = conn2.Tx
	} else {
		repo.db = gorm.Connection.Db
	}

	ommitedFields := []string{"Id", "CreatedAt", "DeletedAt"}

	if result := repo.db.Omit(ommitedFields...).Save(&ent); result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *JobsHierarchiesGorm) Delete(entity ent.JobHierarchy, conn repo.IConnection) error {

	if conn2, ok := conn.(*gorm.ConnectionGorm); ok && conn2 != nil {
		repo.db = conn2.Tx
	} else {
		repo.db = gorm.Connection.Db
	}

	if result := repo.db.Debug().
		Where("uuid = ?", entity.GetID()).
		Delete(&entities.Job{}); result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *JobsHierarchiesGorm) DeleteMany(entitiesToDelete []ent.JobHierarchy, conn repo.IConnection) error {
	var uuidsToDelete []uuid.UUID

	for _, item := range entitiesToDelete {
		uuidsToDelete = append(uuidsToDelete, item.GetID())
	}

	if conn2, ok := conn.(*gorm.ConnectionGorm); ok && conn2 != nil {
		repo.db = conn2.Tx
	} else {
		repo.db = gorm.Connection.Db
	}

	if result := repo.db.Where("uuid IN(?)", uuidsToDelete).Delete(&entities.JobHierarchy{}); result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *JobsHierarchiesGorm) DeleteByJob(entity ent.Job, conn repo.IConnection) error {

	if conn2, ok := conn.(*gorm.ConnectionGorm); ok && conn2 != nil {
		repo.db = conn2.Tx
	} else {
		repo.db = gorm.Connection.Db
	}

	if result := repo.db.Raw(
		`   UPDATE public.jobs_closure jc
          set deleted_at = now()
          FROM public.jobs j
          where j.id = jc.id_job and j.uuid = ?
    `,
		entity.GetID(),
	).Scan(nil); result.Error != nil {
		return result.Error
	}

	return nil
}

func NewJobsHiearchyRepositoryGorm() repositories.IJobsHierarchiesRepository {
	return &JobsHierarchiesGorm{}
}
