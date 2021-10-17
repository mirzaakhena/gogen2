package entity

import (
  "fmt"
  "github.com/mirzaakhena/gogen2/domain/vo"
  "go/ast"
  "go/token"
)

// ObjValueObject ...
type ObjValueObject struct {
  ValueObjectName vo.Naming
  FieldNames      []string
}

// ObjDataValueObject ...
type ObjDataValueObject struct {
  ValueObjectName string
  FieldNames      []string
}

// NewObjValueObject ...
func NewObjValueObject(entityName string, fieldNames []string) (*ObjValueObject, error) {

  var obj ObjValueObject
  obj.ValueObjectName = vo.Naming(entityName)
  obj.FieldNames = fieldNames

  return &obj, nil
}

// GetData ...
func (o ObjValueObject) GetData() *ObjDataValueObject {
  return &ObjDataValueObject{
    ValueObjectName: o.ValueObjectName.String(),
    FieldNames:      o.FieldNames,
  }
}

// GetValueObjectRootFolderName ...
func (o ObjValueObject) GetValueObjectRootFolderName() string {
  return fmt.Sprintf("domain/entity")
}

// GetValueObjectFileName ...
func (o ObjValueObject) GetValueObjectFileName() string {
  return fmt.Sprintf("%s/%s.go", o.GetValueObjectRootFolderName(), o.ValueObjectName.SnakeCase())
}

// IsValueObjectExist ...
func (o ObjValueObject) IsValueObjectExist() (bool, error) {

  fset := token.NewFileSet()
  exist := IsExist(fset, o.GetValueObjectRootFolderName(), func(file *ast.File, ts *ast.TypeSpec) bool {
    _, ok := ts.Type.(*ast.StructType)
    return ok && ts.Name.String() == o.ValueObjectName.String()
  })

  return exist, nil
}
