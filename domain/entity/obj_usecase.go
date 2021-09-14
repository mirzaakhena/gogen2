package entity

import (
  "fmt"
  "github.com/mirzaakhena/gogen2/application/apperror"
  "github.com/mirzaakhena/gogen2/domain/vo"
)

type ObjUsecase struct {
  PackagePath string
  UsecaseName vo.Naming
}

type ObjDataUsecase struct {
  PackagePath string
  UsecaseName string
}

func NewObjUsecase(usecaseName, packagePath string) (*ObjUsecase, error) {

  if usecaseName == "" {
    return nil, apperror.UsecaseNameMustNotEmpty
  }

  var obj ObjUsecase
  obj.UsecaseName = vo.Naming(usecaseName)
  obj.PackagePath = packagePath

  return &obj, nil
}

func (o ObjUsecase) GetData() *ObjDataUsecase {
  return &ObjDataUsecase{
    PackagePath: o.PackagePath,
    UsecaseName: o.UsecaseName.String(),
  }
}

func (o ObjUsecase) GetRootFolderName() string {
  return fmt.Sprintf("usecase/%s", o.UsecaseName.LowerCase())
}

func (o ObjUsecase) GetInportFileName() string {
  return fmt.Sprintf("%s/inport.go", o.GetRootFolderName())
}

func (o ObjUsecase) GetOutportFileName() string {
  return fmt.Sprintf("%s/outport.go", o.GetRootFolderName())
}

func (o ObjUsecase) GetInteractorFileName() string {
  return fmt.Sprintf("%s/interactor.go", o.GetRootFolderName())
}
