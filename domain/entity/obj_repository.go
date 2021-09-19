package entity

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/mirzaakhena/gogen2/domain/vo"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"golang.org/x/tools/imports"
	"io/ioutil"
	"os"
)

type ObjRepository struct {
	RepositoryName vo.Naming
	ObjEntity      ObjEntity
	ObjUsecase     ObjUsecase
}

type ObjDataRepository struct {
	PackagePath    string
	RepositoryName string
	EntityName     string
	UsecaseName    string
}

func NewObjRepository(repositoryName, entityName, usecaseName string) (*ObjRepository, error) {

	var obj ObjRepository
	obj.RepositoryName = vo.Naming(repositoryName)

	uc, err := NewObjUsecase(usecaseName)
	if err != nil {
		return nil, err
	}

	obj.ObjUsecase = *uc

	et, err := NewObjEntity(entityName)
	if err != nil {
		return nil, err
	}

	obj.ObjEntity = *et

	return &obj, nil

}

func (o ObjRepository) GetData(PackagePath string) *ObjDataRepository {
	return &ObjDataRepository{
		PackagePath:    PackagePath,
		RepositoryName: o.RepositoryName.String(),
		EntityName:     o.ObjEntity.EntityName.String(),
		UsecaseName:    o.ObjUsecase.UsecaseName.String(),
	}
}

func GetRepositoryRootFolderName() string {
	return fmt.Sprintf("domain/repository")
}

func GetRepositoryFileName() string {
	return fmt.Sprintf("%s/repository.go", GetRepositoryRootFolderName())
}

func (o ObjRepository) IsRepoExist() (bool, error) {

	fset := token.NewFileSet()

	pkgs, err := parser.ParseDir(fset, GetRepositoryRootFolderName(), nil, parser.ParseComments)
	if err != nil {
		return false, err
	}

	for _, pkg := range pkgs {
		for _, file := range pkg.Files {

			for _, decl := range file.Decls {

				gen, ok := decl.(*ast.GenDecl)
				if !ok || gen.Tok != token.TYPE {
					continue
				}

				for _, specs := range gen.Specs {

					ts, ok := specs.(*ast.TypeSpec)
					if !ok {
						continue
					}

					if _, ok = ts.Type.(*ast.InterfaceType); !ok {
						continue
					}

					// repo already exist, abort the command
					if ts.Name.String() == o.getRepositoryName() {
						return true, nil
					}
				}
			}
		}
	}

	return false, nil
}

func (o ObjRepository) InjectCode(repoTemplateCode string) ([]byte, error) {

	// reopen the file
	file, err := os.Open(GetRepositoryFileName())
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	var buffer bytes.Buffer
	for scanner.Scan() {
		row := scanner.Text()

		buffer.WriteString(row)
		buffer.WriteString("\n")
	}

	if err := file.Close(); err != nil {
		return nil, err
	}

	// write the template in the end of file
	buffer.WriteString(repoTemplateCode)
	buffer.WriteString("\n")

	return buffer.Bytes(), nil
}

func (o ObjRepository) InjectToOutport() (string, error) {

	fileReadPath := GetOutportFileName(o.ObjUsecase)

	// read the outport file directly
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, fileReadPath, nil, parser.ParseComments)
	if err != nil {
		return "", err
	}

	// assume never injected before
	isAlreadyInjectedBefore := false

	// check for every declaration
	for _, decl := range file.Decls {

		gen, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}

		// focus on type
		if gen.Tok != token.TYPE {
			continue
		}

		for _, specs := range gen.Specs {

			ts, ok := specs.(*ast.TypeSpec)
			if !ok {
				continue
			}

			// focus on interface type
			iFace, ok := ts.Type.(*ast.InterfaceType)
			if !ok {
				continue
			}

			// find the specific outport interface with specific "Standards" name
			if ts.Name.String() != OutportInterfaceName {
				continue
			}

			// trace every method to find something line like `repository.SaveOrderRepo`
			for _, meth := range iFace.Methods.List {

				selType, ok := meth.Type.(*ast.SelectorExpr)

				// if interface already injected then abort the mission
				if ok && selType.Sel.String() == o.getRepositoryName() {
					isAlreadyInjectedBefore = true

					// it is already injected.
					// here we exit from loop for check other spec, but it is the only one spec we have
					break
				}

			}

			// we want to inject it now
			if !isAlreadyInjectedBefore {
				// add new repository to outport interface
				iFace.Methods.List = append(iFace.Methods.List, &ast.Field{
					Type: &ast.SelectorExpr{
						X: &ast.Ident{
							Name: GetPackageName(GetRepositoryRootFolderName()),
						},
						Sel: &ast.Ident{
							Name: o.getRepositoryName(),
						},
					},
				})

				// after injection no need to check anymore
				break
			}

		}
	}

	if !isAlreadyInjectedBefore {

		// TODO who is responsible to write a file? entity or gateway?
		// i prefer to use gateway instead of entity

		// rewrite the outport
		f, err := os.Create(fileReadPath)
		if err := printer.Fprint(f, fset, file); err != nil {
			return "", err
		}
		err = f.Close()
		if err != nil {
			return "", err
		}

		// reformat and import
		newBytes, err := imports.Process(fileReadPath, nil, nil)
		if err != nil {
			return "", err
		}

		if err := ioutil.WriteFile(fileReadPath, newBytes, 0644); err != nil {
			return "", err
		}

	}

	// try to inject to interactor
	//obj.injectToInteractor()

	return fileReadPath, nil

}

func (o ObjRepository) getRepositoryName() string {
	return fmt.Sprintf("%sRepo", o.RepositoryName)
}
