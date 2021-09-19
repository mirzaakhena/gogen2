package entity

import (
	"fmt"
	"github.com/mirzaakhena/gogen2/application/apperror"
	"github.com/mirzaakhena/gogen2/domain/vo"
	"go/ast"
	"go/parser"
	"go/token"
)

type ObjEntity struct {
	EntityName vo.Naming
}

type ObjDataEntity struct {
	EntityName string
}

func NewObjEntity(entityName string) (*ObjEntity, error) {

	if entityName == "" {
		return nil, apperror.EntityNameMustNotEmpty
	}

	var obj ObjEntity
	obj.EntityName = vo.Naming(entityName)

	return &obj, nil
}

func (o ObjEntity) GetData() *ObjDataEntity {
	return &ObjDataEntity{
		EntityName: o.EntityName.String(),
	}
}

func GetEntityRootFolderName() string {
	return fmt.Sprintf("domain/entity")
}

func GetEntityFileName(o ObjEntity) string {
	return fmt.Sprintf("%s/%s.go", GetEntityRootFolderName(), o.EntityName)
}

func (o ObjEntity) IsExist() (bool, error) {

	// TODO who is responsible to read a file? entity or gateway?
	// i prefer to use gateway

	fset := token.NewFileSet()

	pkgs, err := parser.ParseDir(fset, GetEntityRootFolderName(), nil, parser.ParseComments)
	if err != nil {
		return false, err
	}

  // in every package
	for _, pkg := range pkgs {

    // in every files
		for _, file := range pkg.Files {

      // in every declaration like type, func, const
			for _, decl := range file.Decls {

				gen, ok := decl.(*ast.GenDecl)

        // focus only to type
				if !ok || gen.Tok != token.TYPE {
					continue
				}

				for _, specs := range gen.Specs {

					ts, ok := specs.(*ast.TypeSpec)
					if !ok {
						continue
					}

          // focus only to struct
					if _, ok = ts.Type.(*ast.StructType); !ok {
						continue
					}

					// entity already exist, abort the command
					if ts.Name.String() == o.EntityName.String() {
						return true, nil
					}
				}
			}

		}
	}

	return false, nil
}
