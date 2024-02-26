package repositories

import (
	"health/core/application/errors"
	repo "health/core/application/repositories"
	ent "health/core/clients/domain/entities"
	"health/core/clients/domain/repositories"
	"health/core/infra/db/gorm"
	"health/core/infra/db/gorm/entities"
	"math/big"

	valueobjects "health/core/application/value-objects"

	"github.com/google/uuid"
	orm "gorm.io/gorm"
)

type JobsGorm struct {
	db   *orm.DB
	job  *entities.Job
	jobs *[]entities.Job
}

func (repo *JobsGorm) FindByUUID(id uuid.UUID, conn repo.IConnection) (*ent.Job, error) {
	if conn2, ok := conn.(*gorm.ConnectionGorm); ok && conn2 != nil {
		repo.db = conn2.Tx
	} else {
		repo.db = gorm.Connection.Db
	}

	if result := repo.db.First(&repo.job, "uuid = ?", id); result.Error != nil {
		return nil, result.Error
	}

	if repo.job == nil {
		return nil, errors.NewNotFoundError("Could not found error using id")
	}

	uniqueTypeTransactionId := valueobjects.NewUniqueUUID(uuid.NullUUID{UUID: repo.job.Uuid})

	ent, err := ent.NewJob(&repo.job.Name, nil, nil, &uniqueTypeTransactionId)

	if err != nil {
		return nil, err
	}

	ent.SetInternalId(repo.job.Id)

	return ent, nil
}

func (repo *JobsGorm) FindByID(id string, conn repo.IConnection) (*ent.Job, error) {
	parsedId, err := uuid.Parse(id)

	if err != nil {
		return nil, err
	}

	return repo.FindByUUID(parsedId, conn)
}

func (repo *JobsGorm) FindByName(name string, conn repo.IConnection) (*ent.Job, error) {
	if conn2, ok := conn.(*gorm.ConnectionGorm); ok && conn2 != nil {
		repo.db = conn2.Tx
	} else {
		repo.db = gorm.Connection.Db
	}

	if result := repo.db.Where("name ILIKE ?", "%"+name+"%").Limit(1).Find(&repo.job); result.Error != nil {
		return nil, result.Error
	}

	if repo.job == nil {
		return nil, errors.NewNotFoundError("Could not found job entity")
	}

	uniqueTypeTransactionId := valueobjects.NewUniqueUUID(uuid.NullUUID{UUID: repo.job.Uuid})
	ent, err := ent.NewJob(&repo.job.Name, nil, nil, &uniqueTypeTransactionId)

	if err != nil {
		return nil, err
	}

	ent.SetInternalId(repo.job.Id)

	return ent, nil

}

func (repo *JobsGorm) Find(params *repositories.SearchParamJob, conn repo.IConnection) []ent.Job {
	if conn2, ok := conn.(*gorm.ConnectionGorm); ok && conn2 != nil {
		repo.db = conn2.Tx
	} else {
		repo.db = gorm.Connection.Db
	}

	if result := repo.db.Omit("id").Find(&repo.jobs); result.Error != nil {
		return []ent.Job{}
	}

	var entities []ent.Job

	for _, item := range *repo.jobs {
		uniqueId := valueobjects.NewUniqueUUID(uuid.NullUUID{UUID: item.Uuid})

		ent, err := ent.NewJob(
			&repo.job.Name,
			nil, nil,
			&uniqueId,
		)

		if err == nil {
			entities = append(entities, *ent)
		}
	}

	return entities
}

func (repo *JobsGorm) FindAndCount(params *repositories.SearchParamJob, conn repo.IConnection) repositories.IResponseSearchableJobs {
	items := repo.Find(params, conn)

	return repositories.IResponseSearchableJobs{
		Total: *big.NewInt(int64(len(items))),
		Items: items,
	}
}

func (repo *JobsGorm) FindManagers(id uuid.UUID, conn repo.IConnection) (*[]ent.Job, error) {
	if conn2, ok := conn.(*gorm.ConnectionGorm); ok && conn2 != nil {
		repo.db = conn2.Tx
	} else {
		repo.db = gorm.Connection.Db
	}

	if result := repo.db.
		Raw(`
            WITH RECURSIVE subordinates as (
                SELECT distinct(jc.id_manager) as id_job FROM
                    public.jobs_closure jc
                    join public.jobs j
                        on j.id = jc.id_job
                    where jc.deleted_at IS NULL AND
                        j.deleted_at IS NULL AND
                        j.uuid = ?
                UNION ALL
                SELECT DISTINCT(jc2.id_manager) as id_job
                    FROM public.jobs_closure jc2
                    join subordinates s
                        on s.id_job = jc2.id_job
                    where jc2.deleted_at IS NULL
            ) SELECT j3.* from subordinates s
                join public.jobs j3 on j3.id = s.id_job
            `,
			id,
		).
		Scan(&repo.jobs); result.Error != nil {
		return nil, result.Error
	}

	var entities []ent.Job

	for _, item := range *repo.jobs {

		uniqueManagerId := valueobjects.NewUniqueUUID(uuid.NullUUID{UUID: item.Uuid})
		managerEnt, err := ent.NewJob(&item.Name, nil, nil, &uniqueManagerId)

		if err == nil {
			entities = append(entities, *managerEnt)
		}
	}

	return &entities, nil
}

func (repo *JobsGorm) FindManagersByJob(job *ent.Job, tx repo.IConnection) (*ent.Job, error) {
	subordinates, err := repo.FindManagers(job.GetID(), tx)

	if err != nil {
		return nil, err
	}

	job.SetManagers(*subordinates)

	return job, nil
}

func (repo *JobsGorm) FindSubordinates(id uuid.UUID, conn repo.IConnection) (*[]ent.Job, error) {
	if conn2, ok := conn.(*gorm.ConnectionGorm); ok && conn2 != nil {
		repo.db = conn2.Tx
	} else {
		repo.db = gorm.Connection.Db
	}

	if result := repo.db.
		Raw(
			`WITH RECURSIVE subordinates as (
            SELECT distinct(jc.id_job) as id_job FROM
                public.jobs_closure jc
                join public.jobs j
                    on j.id = jc.id_manager
                where jc.deleted_at IS NULL and
                    j.deleted_at IS null and
                    j.uuid = ?
            UNION ALL
            SELECT DISTINCT(jc2.id_job) as id_job
                FROM public.jobs_closure jc2
                join subordinates s
                    on s.id_job = jc2.id_manager
                where jc2.deleted_at IS null
        ) SELECT j3.* from subordinates s
            join public.jobs j3 on j3.id = s.id_job
        `,
			id,
		).
		Scan(&repo.jobs); result.Error != nil {
		return nil, result.Error
	}

	var entities []ent.Job

	if repo.jobs != nil {
		for _, item := range *repo.jobs {
			uniqueManagerId := valueobjects.NewUniqueUUID(uuid.NullUUID{UUID: item.Uuid})
			managerEnt, err := ent.NewJob(&item.Name, nil, nil, &uniqueManagerId)

			if err == nil {
				entities = append(entities, *managerEnt)
			}
		}

	}

	return &entities, nil
}

func (repo *JobsGorm) FindSubordinatesByJob(job *ent.Job, tx repo.IConnection) (*ent.Job, error) {
	subordinates, err := repo.FindSubordinates(job.GetID(), tx)

	if err != nil {
		return nil, err
	}

	job.SetSubordinates(*subordinates)

	return job, nil
}

func (repo *JobsGorm) Create(entity *ent.Job, conn repo.IConnection) error {

	ent := entities.Job{
		Uuid:      entity.GetID(),
		Name:      entity.GetName(),
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
	entity.SetInternalId(ent.Id)

	return nil
}

func (repo *JobsGorm) CreateMany(items []ent.Job, conn repo.IConnection) error {

	var itemsEnt []entities.Job

	for _, item := range items {

		ent := entities.Job{
			Uuid:      item.GetID(),
			Name:      item.GetName(),
			CreatedAt: item.CreatedAt,
		}

		itemsEnt = append(itemsEnt, ent)
	}

	ommitFields := []string{"Id", "DeletedAt", "UpdatedAt"}

	if conn2, ok := conn.(*gorm.ConnectionGorm); ok && conn2 != nil {
		repo.db = conn2.Tx
	} else {
		repo.db = gorm.Connection.Db
	}

	if result := repo.db.Omit(ommitFields...).Create(&itemsEnt); result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *JobsGorm) Update(entity ent.Job, conn repo.IConnection) error {

	ent := entities.Job{
		Uuid:      entity.GetID(),
		Name:      entity.GetName(),
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
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

func (repo *JobsGorm) Delete(entity ent.Job, conn repo.IConnection) error {

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

func (repo *JobsGorm) DeleteMany(entitiesToDelete []ent.Job, conn repo.IConnection) error {
	var uuidsToDelete []uuid.UUID

	for _, item := range entitiesToDelete {
		uuidsToDelete = append(uuidsToDelete, item.GetID())
	}

	if conn2, ok := conn.(*gorm.ConnectionGorm); ok && conn2 != nil {
		repo.db = conn2.Tx
	} else {
		repo.db = gorm.Connection.Db
	}

	if result := repo.db.Where("uuid IN(?)", uuidsToDelete).Delete(&entities.Job{}); result.Error != nil {
		return result.Error
	}

	return nil
}

func NewJobsRepositoryGorm() repositories.IJobsRepository {
	return &JobsGorm{}
}
