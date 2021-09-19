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
	"strings"
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

func (o ObjRepository) InjectToOutport() error {

	fileReadPath := GetOutportFileName(o.ObjUsecase)

	// read the outport file directly
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, fileReadPath, nil, parser.ParseComments)
	if err != nil {
		return err
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
			return err
		}
		err = f.Close()
		if err != nil {
			return err
		}

		// reformat and import
		newBytes, err := imports.Process(fileReadPath, nil, nil)
		if err != nil {
			return err
		}

		if err := ioutil.WriteFile(fileReadPath, newBytes, 0644); err != nil {
			return err
		}

	}

	// try to inject to interactor
	//obj.injectToInteractor()

	return nil

}

const injectedCodeLocation = "//!"

func (o ObjRepository) InjectToInteractor(injectedCode string) ([]byte, error) {

	existingFile := GetInteractorFileName(o.ObjUsecase)

	// open interactor file
	file, err := os.Open(existingFile)
	if err != nil {
		return nil, err
	}

	//// check the repo name and return specific template
	//constTemplateCode, err := obj.prepareInteractorTemplate()
	//if err != nil {
	//	return err
	//}

	needToInject := false

	scanner := bufio.NewScanner(file)
	var buffer bytes.Buffer
	for scanner.Scan() {
		row := scanner.Text()

		// check the injected code in interactor
		if strings.TrimSpace(row) == injectedCodeLocation {

			needToInject = true

			//// we need to provide an error
			//InitiateError()

			// inject code
			buffer.WriteString(injectedCode)
			buffer.WriteString("\n")

			continue
		}

		buffer.WriteString(row)
		buffer.WriteString("\n")
	}

	// if no injected marker found, then abort the next step
	if !needToInject {
		return nil, nil
	}

	if err := file.Close(); err != nil {
		return nil, err
	}

	// rewrite the file
	if err := ioutil.WriteFile(existingFile, buffer.Bytes(), 0644); err != nil {
		return nil, err
	}

	//fset := token.NewFileSet()
	//node, err := parser.ParseFile(fset, existingFile, nil, parser.ParseComments)
	//if err != nil {
	//	return nil,  err
	//}
	//
	//// reformat the code
	//var newBuf bytes.Buffer
	//err = format.Node(&newBuf, fset, node)
	//if err != nil {
	//	return nil, err
	//}
	//
	//// rewrite again
	//if err := ioutil.WriteFile(existingFile, newBuf.Bytes(), 0644); err != nil {
	//	return nil, err
	//}

	return buffer.Bytes(), nil
}

func (o ObjRepository) getRepositoryName() string {
	return fmt.Sprintf("%sRepo", o.RepositoryName)
}
