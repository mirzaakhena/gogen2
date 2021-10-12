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
  DriverName     string
}

// ObjDataController ...
type ObjDataController struct {
  PackagePath    string
  UsecaseName    string
  ControllerName string
}

func NewObjController(controllerName, driverName string) (*ObjController, error) {

  if controllerName == "" {
    return nil, apperror.ControllerNameMustNotEmpty
  }

  var obj ObjController
  obj.ControllerName = vo.Naming(controllerName)
  obj.DriverName = driverName

  return &obj, nil
}

// GetData ...
func (o ObjController) GetData(PackagePath string, ou ObjUsecase) *ObjDataController {
  return &ObjDataController{
    PackagePath:    PackagePath,
    ControllerName: o.ControllerName.String(),
    UsecaseName:    ou.UsecaseName.String(),
  }
}

// GetControllerRootFolderName ...
func (o ObjController) GetControllerRootFolderName() string {
  return fmt.Sprintf("controller/%s", o.ControllerName.LowerCase())
}

// GetControllerInterfaceFile ...
func (o ObjController) GetControllerInterfaceFile() string {
  return fmt.Sprintf("controller/controller.go")
}

// GetControllerResponseFileName ...
func (o ObjController) GetControllerResponseFileName() string {
  return fmt.Sprintf("%s/response.go", o.GetControllerRootFolderName())
}

// GetControllerInterceptorFileName ...
func (o ObjController) GetControllerInterceptorFileName() string {
  return fmt.Sprintf("%s/interceptor.go", o.GetControllerRootFolderName())
}

// GetControllerRouterFileName ...
func (o ObjController) GetControllerRouterFileName() string {
  return fmt.Sprintf("%s/router.go", o.GetControllerRootFolderName())
}

// GetControllerHandlerFileName ...
func (o ObjController) GetControllerHandlerFileName(ou ObjUsecase) string {
  return fmt.Sprintf("%s/handler_%s.go", o.GetControllerRootFolderName(), ou.UsecaseName.LowerCase())
}

func (o ObjController) InjectInportToStruct(templateWithData string) ([]byte, error) {

  inportLine, err := o.getInportLine()
  if err != nil {
    return nil, err
  }

  controllerFile := o.GetControllerRouterFileName()

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

  controllerFile := o.GetControllerRouterFileName()

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

  return buffer.Bytes(), nil

}

func (o ObjController) getInportLine() (int, error) {

  controllerFile := o.GetControllerRouterFileName()

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

  controllerFile := o.GetControllerRouterFileName()

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
