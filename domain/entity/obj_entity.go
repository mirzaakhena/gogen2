package entity

import (
	"fmt"
	"github.com/mirzaakhena/gogen2/application/apperror"
	"github.com/mirzaakhena/gogen2/domain/vo"
	"go/ast"
	"go/token"
)

// ObjEntity ...
type ObjEntity struct {
	EntityName vo.Naming
}

// ObjDataEntity ...
type ObjDataEntity struct {
	EntityName string
}

// NewObjEntity ...
func NewObjEntity(entityName string) (*ObjEntity, error) {

	if entityName == "" {
		return nil, apperror.EntityNameMustNotEmpty
	}

	var obj ObjEntity
	obj.EntityName = vo.Naming(entityName)

	return &obj, nil
}

// GetData ...
func (o ObjEntity) GetData() *ObjDataEntity {
	return &ObjDataEntity{
		EntityName: o.EntityName.String(),
	}
}

// GetEntityRootFolderName ...
func (o ObjEntity) GetEntityRootFolderName() string {
	return fmt.Sprintf("domain/entity")
}

// GetEntityFileName ...
func (o ObjEntity) GetEntityFileName() string {
	return fmt.Sprintf("%s/%s.go", o.GetEntityRootFolderName(), o.EntityName.SnakeCase())
}

// IsEntityExist ...
func (o ObjEntity) IsEntityExist() (bool, error) {

	fset := token.NewFileSet()
	exist := IsExist(fset, o.GetEntityRootFolderName(), func(file *ast.File, ts *ast.TypeSpec) bool {
		_, ok := ts.Type.(*ast.StructType)
		return ok && ts.Name.String() == o.EntityName.String()
	})

	return exist, nil
}
