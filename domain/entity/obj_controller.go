package entity

import (
  "bufio"
  "bytes"
  "fmt"
  "github.com/mirzaakhena/gogen2/application/apperror"
  "github.com/mirzaakhena/gogen2/domain/vo"
  "go/ast"
  "go/parser"
  "go/token"
  "os"
)

type ObjController struct {
  ControllerName vo.Naming
  ObjUsecase     ObjUsecase
}

// ObjDataController ...
type ObjDataController struct {
  PackagePath    string
  UsecaseName    string
  ControllerName string
}

func NewObjController(controllerName string, objUsecase ObjUsecase) (*ObjController, error) {

  if controllerName == "" {
    return nil, apperror.ControllerNameMustNotEmpty
  }

  var obj ObjController
  obj.ControllerName = vo.Naming(controllerName)
  obj.ObjUsecase = objUsecase

  return &obj, nil
}

// GetData ...
func (o ObjController) GetData(PackagePath string) *ObjDataController {
  return &ObjDataController{
    PackagePath:    PackagePath,
    ControllerName: o.ControllerName.String(),
    UsecaseName:    o.ObjUsecase.UsecaseName.String(),
  }
}

// GetControllerRootFolderName ...
func GetControllerRootFolderName(o ObjController) string {
  return fmt.Sprintf("controller/%s", o.ControllerName.LowerCase())
}

// GetControllerResponseFileName ...
func GetControllerResponseFileName(o ObjController) string {
  return fmt.Sprintf("%s/response.go", GetControllerRootFolderName(o))
}

// GetControllerInterceptorFileName ...
func GetControllerInterceptorFileName(o ObjController) string {
  return fmt.Sprintf("%s/interceptor.go", GetControllerRootFolderName(o))
}

// GetControllerRouterFileName ...
func GetControllerRouterFileName(o ObjController) string {
  return fmt.Sprintf("%s/router.go", GetControllerRootFolderName(o))
}

// GetControllerHandlerFileName ...
func GetControllerHandlerFileName(o ObjController) string {
  return fmt.Sprintf("%s/handler_%s.go", GetControllerRootFolderName(o), o.ObjUsecase.UsecaseName.LowerCase())
}

func (o ObjController) InjectInportToStruct(templateWithData string) ([]byte, error) {

  inportLine, err := o.getInportLine()
  if err != nil {
    return nil, err
  }

  controllerFile := GetControllerRouterFileName(o)

  file, err := os.Open(controllerFile)
  if err != nil {
    return nil, err
  }
  defer func() {
    if err := file.Close(); err != nil {
      return
    }
  }()

  scanner := bufio.NewScanner(file)
  var buffer bytes.Buffer
  line := 0
  for scanner.Scan() {
    row := scanner.Text()

    if line == inportLine-1 {
      buffer.WriteString(templateWithData)
      buffer.WriteString("\n")
    }

    buffer.WriteString(row)
    buffer.WriteString("\n")
    line++
  }

  return buffer.Bytes(), nil
}

func (o ObjController) InjectRouterBind(templateWithData string) ([]byte, error) {
  {

    controllerFile := GetControllerRouterFileName(o)

    routerLine, err := o.getBindRouterLine()
    if err != nil {
      return nil, err
    }

    //templateCode, err := util.PrintTemplate(templates.ControllerBindRouterGinFile, obj)
    //if err != nil {
    //  return err
    //}

    file, err := os.Open(controllerFile)
    if err != nil {
      return nil, err
    }
    defer func() {
      if err := file.Close(); err != nil {
        return
      }
    }()

    scanner := bufio.NewScanner(file)
    var buffer bytes.Buffer
    line := 0
    for scanner.Scan() {
      row := scanner.Text()

      if line == routerLine-1 {
        buffer.WriteString(templateWithData)
        buffer.WriteString("\n")
      }

      buffer.WriteString(row)
      buffer.WriteString("\n")
      line++
    }

    //// reformat and do import
    //newBytes, err := imports.Process(controllerFile, buffer.Bytes(), nil)
    //if err != nil {
    //  return err
    //}
    //
    //if err := ioutil.WriteFile(controllerFile, newBytes, 0644); err != nil {
    //  return err
    //}

    return buffer.Bytes(), nil

  }
}

func (o ObjController) getInportLine() (int, error) {

  controllerFile := GetControllerRouterFileName(o)

  inportLine := 0
  fset := token.NewFileSet()
  astFile, err := parser.ParseFile(fset, controllerFile, nil, parser.ParseComments)
  if err != nil {
    return 0, err
  }

  // loop the outport for imports
  for _, decl := range astFile.Decls {

    if gen, ok := decl.(*ast.GenDecl); ok {

      if gen.Tok != token.TYPE {
        continue
      }

      for _, specs := range gen.Specs {

        ts, ok := specs.(*ast.TypeSpec)
        if !ok {
          continue
        }

        if iStruct, ok := ts.Type.(*ast.StructType); ok {

          // check the specific struct name
          if ts.Name.String() != "Controller" {
            continue
          }

          inportLine = fset.Position(iStruct.Fields.Closing).Line
          return inportLine, nil
        }

      }

    }

  }

  return 0, fmt.Errorf(" Controller struct not found")

}

func (o ObjController) getBindRouterLine() (int, error) {

  controllerFile := GetControllerRouterFileName(o)

  fset := token.NewFileSet()
  astFile, err := parser.ParseFile(fset, controllerFile, nil, parser.ParseComments)
  if err != nil {
    return 0, err
  }
  routerLine := 0
  for _, decl := range astFile.Decls {

    if gen, ok := decl.(*ast.FuncDecl); ok {

      if gen.Recv == nil {
        continue
      }

      startExp, ok := gen.Recv.List[0].Type.(*ast.StarExpr)
      if !ok {
        continue
      }

      if startExp.X.(*ast.Ident).String() != "Controller" {
        continue
      }

      if gen.Name.String() != "RegisterRouter" {
        continue
      }

      routerLine = fset.Position(gen.Body.Rbrace).Line
      return routerLine, nil
    }

  }
  return 0, fmt.Errorf("register router Not found")
}
