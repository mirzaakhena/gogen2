package entity

import (
  "bufio"
  "bytes"
  "fmt"
  "github.com/mirzaakhena/gogen2/domain/vo"
  "go/ast"
  "go/parser"
  "go/token"
  "golang.org/x/tools/imports"
  "io/ioutil"
  "os"
)

type ObjRepository struct {
  PackagePath    string
  RepositoryName vo.Naming
  EntityName     vo.Naming
  UsecaseName    vo.Naming
}

type ObjDataRepository struct {
  PackagePath    string
  RepositoryName string
  EntityName     string
  UsecaseName    string
}

func NewObjRepository(repositoryName, entityName, usecaseName, packagePath string) (*ObjRepository, error) {

  var obj ObjRepository
  obj.RepositoryName = vo.Naming(repositoryName)
  obj.EntityName = vo.Naming(entityName)
  obj.UsecaseName = vo.Naming(usecaseName)
  obj.PackagePath = packagePath

  return &obj, nil

}

func (o ObjRepository) GetRootFolderName() string {
  return fmt.Sprintf("domain/repository")
}

func (o ObjRepository) GetMainFileName() string {
  return fmt.Sprintf("%s/repository.go", o.GetRootFolderName())
}

func (o ObjRepository) IsRepoExist() (bool, error) {

  fset := token.NewFileSet()

  pkgs, err := parser.ParseDir(fset, o.GetRootFolderName(), nil, parser.ParseComments)
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
          if ts.Name.String() == fmt.Sprintf("%sRepo", o.RepositoryName) {
            return true, nil
          }
        }
      }
    }
  }

  return false, nil
}

func (o ObjRepository) InjectCode(repoTemplateCode string) error {

  // reopen the file
  file, err := os.Open(o.GetMainFileName())
  if err != nil {
    return err
  }

  scanner := bufio.NewScanner(file)
  var buffer bytes.Buffer
  for scanner.Scan() {
    row := scanner.Text()

    buffer.WriteString(row)
    buffer.WriteString("\n")
  }

  if err := file.Close(); err != nil {
    return err
  }

  // write the template in the end of file
  buffer.WriteString(repoTemplateCode)
  buffer.WriteString("\n")

  // reformat and do import
  newBytes, err := imports.Process(o.GetMainFileName(), buffer.Bytes(), nil)
  if err != nil {
    return err
  }

  if err := ioutil.WriteFile(o.GetMainFileName(), newBytes, 0644); err != nil {
    return err
  }

  return nil
}
