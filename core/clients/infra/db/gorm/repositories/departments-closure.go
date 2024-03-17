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

type DepartmentsHierarchiesGorm struct {
	db                 *orm.DB
	department         *entities.DepartmentHierarchy
	departmentsClosure *[]entities.DepartmentHierarchy
}

func (repo *DepartmentsHierarchiesGorm) FindByUUID(id uuid.UUID, conn repo.IConnection) (*ent.DepartmentHierarchy, error) {
	if conn2, ok := conn.(*gorm.ConnectionGorm); ok && conn2 != nil {
		repo.db = conn2.Tx
	} else {
		repo.db = gorm.Connection.Db
	}

	if result := repo.db.Omit("Id").InnerJoins("Job").First(&repo.department, "jobs_closure.uuid = ?", id); result.Error != nil {
		return nil, result.Error
	}

	if repo.department == nil {
		return nil, errors.NewNotFoundError("Could not found error using id")
	}

	jobGorm := repo.department.Department
	uniqueIddepartment := valueobjects.NewUniqueUUID(uuid.NullUUID{UUID: jobGorm.Uuid})

	job, _ := ent.NewDepartment(&jobGorm.Name, nil, nil, &uniqueIddepartment)

	uniqueTypeTransactionId := valueobjects.NewUniqueUUID(uuid.NullUUID{UUID: repo.department.Uuid})

	return ent.NewDepartmentHierarchy(job, nil, &uniqueTypeTransactionId)
}

func (repo *DepartmentsHierarchiesGorm) FindByID(id string, conn repo.IConnection) (*ent.DepartmentHierarchy, error) {
	parsedId, err := uuid.Parse(id)

	if err != nil {
		return nil, err
	}

	return repo.FindByUUID(parsedId, conn)
}

func (repo *DepartmentsHierarchiesGorm) Find(params *repositories.SearchParamDepartmentHierarchy, conn repo.IConnection) []ent.DepartmentHierarchy {
	if conn2, ok := conn.(*gorm.ConnectionGorm); ok && conn2 != nil {
		repo.db = conn2.Tx
	} else {
		repo.db = gorm.Connection.Db
	}

	if result := repo.db.Omit("id").InnerJoins("Job").Find(&repo.departmentsClosure); result.Error != nil {
		return []ent.DepartmentHierarchy{}
	}

	var entities []ent.DepartmentHierarchy

	for _, item := range *repo.departmentsClosure {
		department := item.Department

		uniqueJobId := valueobjects.NewUniqueUUID(uuid.NullUUID{UUID: department.Uuid})
		jobEntity, _ := ent.NewDepartment(&department.Name, nil, nil, &uniqueJobId)

		uniqueId := valueobjects.NewUniqueUUID(uuid.NullUUID{UUID: item.Uuid})

		ent, err := ent.NewDepartmentHierarchy(
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

func (repo *DepartmentsHierarchiesGorm) FindAndCount(params *repositories.SearchParamDepartmentHierarchy, conn repo.IConnection) repositories.IResponseSearchableDepartmentsHirearchy {
	items := repo.Find(params, conn)

	return repositories.IResponseSearchableDepartmentsHirearchy{
		Total: *big.NewInt(int64(len(items))),
		Items: items,
	}
}

func (repo *DepartmentsHierarchiesGorm) Create(entity *ent.DepartmentHierarchy, conn repo.IConnection) error {

	department := entity.GetDepartment()
	manager := entity.GetManager()

	var idManager *int

	if manager != nil {
		internalId := manager.GetInternalId()
		idManager = &internalId
	}

	ent := entities.DepartmentHierarchy{
		Uuid:         entity.GetID(),
		IdDepartment: department.GetInternalId(),
		IdManager:    idManager,
		CreatedAt:    entity.CreatedAt,
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

func (repo *DepartmentsHierarchiesGorm) CreateMany(newEntities []ent.DepartmentHierarchy, conn repo.IConnection) error {

	var hierarchies []entities.DepartmentHierarchy

	for _, item := range newEntities {
		department := item.GetDepartment()
		manager := item.GetManager()
		var idManager *int

		if manager != nil {
			id := manager.GetInternalId()
			idManager = &id
		}

		hierarchies = append(hierarchies, entities.DepartmentHierarchy{
			Uuid:         item.GetID(),
			IdDepartment: department.GetInternalId(),
			IdManager:    idManager,
			CreatedAt:    item.CreatedAt,
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

func (repo *DepartmentsHierarchiesGorm) Update(entity ent.DepartmentHierarchy, conn repo.IConnection) error {

	department := entity.GetDepartment()
	manager := entity.GetManager()

	var idManager *int

	if manager != nil {
		internalId := manager.GetInternalId()
		idManager = &internalId
	}

	ent := entities.DepartmentHierarchy{
		Uuid:         entity.GetID(),
		IdDepartment: department.GetInternalId(),
		IdManager:    idManager,
		CreatedAt:    entity.CreatedAt,
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

func (repo *DepartmentsHierarchiesGorm) Delete(entity ent.DepartmentHierarchy, conn repo.IConnection) error {

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

func (repo *DepartmentsHierarchiesGorm) DeleteMany(entitiesToDelete []ent.DepartmentHierarchy, conn repo.IConnection) error {
	var uuidsToDelete []uuid.UUID

	for _, item := range entitiesToDelete {
		uuidsToDelete = append(uuidsToDelete, item.GetID())
	}

	if conn2, ok := conn.(*gorm.ConnectionGorm); ok && conn2 != nil {
		repo.db = conn2.Tx
	} else {
		repo.db = gorm.Connection.Db
	}

	if result := repo.db.Where("uuid IN(?)", uuidsToDelete).Delete(&entities.DepartmentHierarchy{}); result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *DepartmentsHierarchiesGorm) DeleteByDepartment(entity ent.Department, conn repo.IConnection) error {

	if conn2, ok := conn.(*gorm.ConnectionGorm); ok && conn2 != nil {
		repo.db = conn2.Tx
	} else {
		repo.db = gorm.Connection.Db
	}

	if result := repo.db.Raw(
		`   UPDATE public.departments_closure jc
          set deleted_at = now()
          FROM public.departments j
          where j.id = jc.id_department and j.uuid = ?
    `,
		entity.GetID(),
	).Scan(nil); result.Error != nil {
		return result.Error
	}

	return nil
}

func NewDepartmentsHiearchyRepositoryGorm() repositories.IDepartmentsHierarchiesRepository {
	return &DepartmentsHierarchiesGorm{}
}
