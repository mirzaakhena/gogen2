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

func NewObjEntity(usecaseName string) (*ObjEntity, error) {

	if usecaseName == "" {
		return nil, apperror.EntityNameMustNotEmpty
	}

	var obj ObjEntity
	obj.EntityName = vo.Naming(usecaseName)

	return &obj, nil
}

func (o ObjEntity) GetData() *ObjDataEntity {
	return &ObjDataEntity{
		EntityName: o.EntityName.String(),
	}
}

func (o ObjEntity) GetRootFolderName() string {
	return fmt.Sprintf("domain/entity")
}

func (o ObjEntity) GetEntityFileName() string {
	return fmt.Sprintf("%s/%s.go", o.GetRootFolderName(), o.EntityName)
}

func (o ObjEntity) IsExist() (bool, error) {

	fset := token.NewFileSet()

	pkgs, err := parser.ParseDir(fset, o.GetRootFolderName(), nil, parser.ParseComments)
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
					if ts.Name.String() == fmt.Sprintf("%s", o.EntityName) {
						return true, nil
					}
				}
			}

		}
	}

	return false, nil
}
