package entity

import (
  "fmt"
  "github.com/mirzaakhena/gogen2/domain/vo"
  "strings"
)

type ObjTesting struct {
  PackagePath string
  UsecaseName vo.Naming
  TestName    vo.Naming
  Methods     vo.OutportMethods
}

type ObjDataTesting struct {
  PackagePath string
  UsecaseName string
  TestName    string
  Methods     vo.OutportMethods
}

func NewObjTesting(testName, usecaseName, packagePath string) (*ObjTesting, error) {

  var obj ObjTesting
  obj.TestName = vo.Naming(testName)
  obj.UsecaseName = vo.Naming(usecaseName)
  obj.PackagePath = packagePath

  err := obj.Methods.ReadOutport(usecaseName, packagePath)
  if err != nil {
    return nil, err
  }

  // Little hacky for replace the context.Context{} with ctx variable
  for _, m := range obj.Methods {
    if strings.HasPrefix(strings.TrimSpace(m.DefaultParamVal), `context.Context{}`) {
      m.DefaultParamVal = strings.ReplaceAll(m.DefaultParamVal, `context.Context{}`, "ctx")
    }
  }

  return &obj, nil
}

func (o ObjTesting) GetData() *ObjDataTesting {
  return &ObjDataTesting{
    PackagePath: o.PackagePath,
    UsecaseName: o.UsecaseName.String(),
    TestName:    o.TestName.LowerCase(),
    Methods:     o.Methods,
  }
}

func (o ObjTesting) GetTestFileName() string {
  return fmt.Sprintf("usecase/%s/testcase_%s_test.go", o.UsecaseName.LowerCase(), o.TestName.LowerCase())
}