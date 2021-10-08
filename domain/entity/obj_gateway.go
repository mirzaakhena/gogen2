package entity

import (
  "fmt"
  "github.com/mirzaakhena/gogen2/application/apperror"
  "github.com/mirzaakhena/gogen2/domain/vo"
  "go/ast"
  "go/parser"
  "go/token"
  "os"
  "strings"
)

const gatewayStructName = "gateway"

// ObjGateway  depend on (which) usecase that want to be tested
type ObjGateway struct {
  GatewayName vo.Naming
}

// ObjDataGateway  ...
type ObjDataGateway struct {
  PackagePath string
  GatewayName string
  Methods     vo.OutportMethods
}

// NewObjGateway   ...
func NewObjGateway(gatewayName string) (*ObjGateway, error) {

  if gatewayName == "" {
    return nil, apperror.GatewayNameMustNotEmpty
  }

  var obj ObjGateway
  obj.GatewayName = vo.Naming(gatewayName)

  return &obj, nil
}

// GetData ...
func (o ObjGateway) GetData(PackagePath string, outportMethods vo.OutportMethods) *ObjDataGateway {
  return &ObjDataGateway{
    PackagePath: PackagePath,
    GatewayName: o.GatewayName.LowerCase(),
    Methods:     outportMethods,
  }
}

// GetGatewayRootFolderName ...
func GetGatewayRootFolderName(o ObjGateway) string {
  return fmt.Sprintf("gateway/%s", o.GatewayName.LowerCase())
}

// GetGatewayFileName ...
func GetGatewayFileName(o ObjGateway) string {
  return fmt.Sprintf("%s/implementation.go", GetGatewayRootFolderName(o))
}

// GetGatewayStructName ...
func GetGatewayStructName(o ObjGateway) string {
  return fmt.Sprintf("%sGateway", o.GatewayName.CamelCase())
}

func (o ObjGateway) InjectToGateway(injectedCode string) ([]byte, error) {
  return InjectCodeAtTheEndOfFile(GetGatewayFileName(o), injectedCode)
}

func FindGatewayByName(gatewayName string) (*ObjGateway, error) {
  folderName := fmt.Sprintf("gateway/%s", strings.ToLower(gatewayName))

  fset := token.NewFileSet()
  pkgs, err := parser.ParseDir(fset, folderName, nil, parser.ParseComments)
  if err != nil {
    return nil, err
  }

  for _, pkg := range pkgs {

    // read file by file
    for _, file := range pkg.Files {

      // in every declaration like type, func, const
      for _, decl := range file.Decls {

        // focus only to type
        gen, ok := decl.(*ast.GenDecl)
        if !ok || gen.Tok != token.TYPE {
          continue
        }

        for _, specs := range gen.Specs {

          ts, ok := specs.(*ast.TypeSpec)
          if !ok {
            continue
          }

          if _, ok := ts.Type.(*ast.StructType); ok {

            // check the specific struct name
            if ts.Name.String() != gatewayStructName {
              continue
            }

            return NewObjGateway(pkg.Name)
            //inportLine = fset.Position(iStruct.Fields.Closing).Line
            //return inportLine, nil
          }
        }

      }

    }

  }

  return nil, err
}

func FindAllObjGateway() ([]*ObjGateway, error) {

  dir, err := os.ReadDir("gateway")
  if err != nil {
    return nil, err
  }

  for _, d := range dir {
    name, err := FindGatewayByName(d.Name())
    if err != nil {
      return nil, err
    }

    fmt.Printf(">>>>>>>> %s\n", name)
  }

	return nil, nil
}


