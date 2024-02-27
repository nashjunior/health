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

type DepartmentsGorm struct {
	db          *orm.DB
	department  *entities.Department
	departments *[]entities.Department
}

func (repo *DepartmentsGorm) FindByUUID(id uuid.UUID, conn repo.IConnection) (*ent.Department, error) {
	if conn2, ok := conn.(*gorm.ConnectionGorm); ok && conn2 != nil {
		repo.db = conn2.Tx
	} else {
		repo.db = gorm.Connection.Db
	}

	if result := repo.db.First(&repo.department, "uuid = ?", id); result.Error != nil {
		return nil, result.Error
	}

	if repo.department == nil {
		return nil, errors.NewNotFoundError("Could not found error using id")
	}

	uniqueTypeTransactionId := valueobjects.NewUniqueUUID(uuid.NullUUID{UUID: repo.department.Uuid})

	ent, err := ent.NewDepartment(&repo.department.Name, nil, nil, &uniqueTypeTransactionId)

	if err != nil {
		return nil, err
	}

	ent.SetInternalId(repo.department.Id)

	return ent, nil
}

func (repo *DepartmentsGorm) FindByID(id string, conn repo.IConnection) (*ent.Department, error) {
	parsedId, err := uuid.Parse(id)

	if err != nil {
		return nil, err
	}

	return repo.FindByUUID(parsedId, conn)
}

func (repo *DepartmentsGorm) FindByName(name string, conn repo.IConnection) (*ent.Department, error) {
	if conn2, ok := conn.(*gorm.ConnectionGorm); ok && conn2 != nil {
		repo.db = conn2.Tx
	} else {
		repo.db = gorm.Connection.Db
	}

	if result := repo.db.Where("name ILIKE ?", "%"+name+"%").Limit(1).Find(&repo.department); result.Error != nil {
		return nil, result.Error
	}

	if repo.department == nil {
		return nil, errors.NewNotFoundError("Could not found job entity")
	}

	uniqueTypeTransactionId := valueobjects.NewUniqueUUID(uuid.NullUUID{UUID: repo.department.Uuid})
	ent, err := ent.NewDepartment(&repo.department.Name, nil, nil, &uniqueTypeTransactionId)

	if err != nil {
		return nil, err
	}

	ent.SetInternalId(repo.department.Id)

	return ent, nil

}

func (repo *DepartmentsGorm) Find(params *repositories.SearchParamDepartment, conn repo.IConnection) []ent.Department {
	if conn2, ok := conn.(*gorm.ConnectionGorm); ok && conn2 != nil {
		repo.db = conn2.Tx
	} else {
		repo.db = gorm.Connection.Db
	}

	if result := repo.db.Omit("id").Find(&repo.departments); result.Error != nil {
		return []ent.Department{}
	}

	var entities []ent.Department

	for _, item := range *repo.departments {
		uniqueId := valueobjects.NewUniqueUUID(uuid.NullUUID{UUID: item.Uuid})

		ent, err := ent.NewDepartment(
			&repo.department.Name,
			nil, nil,
			&uniqueId,
		)

		if err == nil {
			entities = append(entities, *ent)
		}
	}

	return entities
}

func (repo *DepartmentsGorm) FindAndCount(params *repositories.SearchParamDepartment, conn repo.IConnection) repositories.IResponseSearchableDepartments {
	items := repo.Find(params, conn)

	return repositories.IResponseSearchableDepartments{
		Total: *big.NewInt(int64(len(items))),
		Items: items,
	}
}

func (repo *DepartmentsGorm) FindManagers(id uuid.UUID, conn repo.IConnection) (*[]ent.Department, error) {
	if conn2, ok := conn.(*gorm.ConnectionGorm); ok && conn2 != nil {
		repo.db = conn2.Tx
	} else {
		repo.db = gorm.Connection.Db
	}

	if result := repo.db.
		Raw(`
            WITH RECURSIVE subordinates as (
                SELECT distinct(jc.id_manager) as id_department FROM
                    public.departments_closure jc
                    join public.departments j
                        on j.id = jc.id_department
                    where jc.deleted_at IS NULL AND
                        j.deleted_at IS NULL AND
                        j.uuid = ?
                UNION ALL
                SELECT DISTINCT(jc2.id_manager) as id_department
                    FROM public.departments_closure jc2
                    join subordinates s
                        on s.id_department = jc2.id_department
                    where jc2.deleted_at IS NULL
            ) SELECT j3.* from subordinates s
                join public.departments j3 on j3.id = s.id_department
            `,
			id,
		).
		Scan(&repo.departments); result.Error != nil {
		return nil, result.Error
	}

	var entities []ent.Department

	for _, item := range *repo.departments {

		uniqueManagerId := valueobjects.NewUniqueUUID(uuid.NullUUID{UUID: item.Uuid})
		managerEnt, err := ent.NewDepartment(&item.Name, nil, nil, &uniqueManagerId)

		if err == nil {
			entities = append(entities, *managerEnt)
		}
	}

	return &entities, nil
}

func (repo *DepartmentsGorm) FindManagersByDepartment(department *ent.Department, tx repo.IConnection) (*ent.Department, error) {
	subordinates, err := repo.FindManagers(department.GetID(), tx)

	if err != nil {
		return nil, err
	}

	department.SetManagers(*subordinates)

	return department, nil
}

func (repo *DepartmentsGorm) FindSubordinates(id uuid.UUID, conn repo.IConnection) (*[]ent.Department, error) {
	if conn2, ok := conn.(*gorm.ConnectionGorm); ok && conn2 != nil {
		repo.db = conn2.Tx
	} else {
		repo.db = gorm.Connection.Db
	}

	if result := repo.db.
		Raw(
			`WITH RECURSIVE subordinates as (
            SELECT distinct(jc.id_department) as id_department FROM
                public.departments_closure jc
                join public.departments j
                    on j.id = jc.id_manager
                where jc.deleted_at IS NULL and
                    j.deleted_at IS null and
                    j.uuid = ?
            UNION ALL
            SELECT DISTINCT(jc2.id_department) as id_department
                FROM public.departments_closure jc2
                join subordinates s
                    on s.id_department = jc2.id_manager
                where jc2.deleted_at IS null
        ) SELECT j3.* from subordinates s
            join public.departments j3 on j3.id = s.id_department
        `,
			id,
		).
		Scan(&repo.departments); result.Error != nil {
		return nil, result.Error
	}

	var entities []ent.Department

	if repo.departments != nil {
		for _, item := range *repo.departments {
			uniqueManagerId := valueobjects.NewUniqueUUID(uuid.NullUUID{UUID: item.Uuid})
			managerEnt, err := ent.NewDepartment(&item.Name, nil, nil, &uniqueManagerId)

			if err == nil {
				entities = append(entities, *managerEnt)
			}
		}

	}

	return &entities, nil
}

func (repo *DepartmentsGorm) FindSubordinatesByDepartment(job *ent.Department, tx repo.IConnection) (*ent.Department, error) {
	subordinates, err := repo.FindSubordinates(job.GetID(), tx)

	if err != nil {
		return nil, err
	}

	job.SetSubordinates(*subordinates)

	return job, nil
}

func (repo *DepartmentsGorm) Create(entity *ent.Department, conn repo.IConnection) error {

	ent := entities.Department{
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

func (repo *DepartmentsGorm) CreateMany(items []ent.Department, conn repo.IConnection) error {

	var itemsEnt []entities.Department

	for _, item := range items {

		ent := entities.Department{
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

func (repo *DepartmentsGorm) Update(entity ent.Department, conn repo.IConnection) error {

	ent := entities.Department{
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

func (repo *DepartmentsGorm) Delete(entity ent.Department, conn repo.IConnection) error {

	if conn2, ok := conn.(*gorm.ConnectionGorm); ok && conn2 != nil {
		repo.db = conn2.Tx
	} else {
		repo.db = gorm.Connection.Db
	}

	if result := repo.db.Debug().
		Where("uuid = ?", entity.GetID()).
		Delete(&entities.Department{}); result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *DepartmentsGorm) DeleteMany(entitiesToDelete []ent.Department, conn repo.IConnection) error {
	var uuidsToDelete []uuid.UUID

	for _, item := range entitiesToDelete {
		uuidsToDelete = append(uuidsToDelete, item.GetID())
	}

	if conn2, ok := conn.(*gorm.ConnectionGorm); ok && conn2 != nil {
		repo.db = conn2.Tx
	} else {
		repo.db = gorm.Connection.Db
	}

	if result := repo.db.Where("uuid IN(?)", uuidsToDelete).Delete(&entities.Department{}); result.Error != nil {
		return result.Error
	}

	return nil
}

func NewDepartmentssRepositoryGorm() repositories.IDepartmentsRepository {
	return &DepartmentsGorm{}
}
