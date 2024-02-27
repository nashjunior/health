package entities

import (
	"health/core/application/entities"
	valueobjects "health/core/application/value-objects"

	"github.com/go-playground/validator/v10"
)

var validationDepartmentHierarchy *validator.Validate

type DepartmentHierarchy struct {
	entities.Entity

	id *int

	department *Department
	manager    *Department
}

func (department *DepartmentHierarchy) GetDepartment() Department {
	return *department.department
}

func (department *DepartmentHierarchy) GetManager() *Department {
	return department.manager
}

func (department *DepartmentHierarchy) Update(manager *Department) {
	if manager != nil {
		department.manager = manager
	}
}

func NewDepartmentHierarchy(
	department *Department,
	manager *Department,
	id *valueobjects.UniqueEntityUUID,
) (*DepartmentHierarchy, error) {
	departmentTree := &DepartmentHierarchy{Entity: entities.NewEntity(id, nil)}

	departmentTree.department = department
	departmentTree.manager = manager
	return departmentTree, nil
}
