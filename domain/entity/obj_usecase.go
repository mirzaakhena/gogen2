package entity

import (
  "fmt"
  "github.com/mirzaakhena/gogen2/application/apperror"
  "github.com/mirzaakhena/gogen2/domain/vo"
)

const (
  OutportInterfaceName = "Outport"
)

type ObjUsecase struct {
  UsecaseName vo.Naming
}

type ObjDataUsecase struct {
  PackagePath string
  UsecaseName string
}

func NewObjUsecase(usecaseName string) (*ObjUsecase, error) {

  if usecaseName == "" {
    return nil, apperror.UsecaseNameMustNotEmpty
  }

  var obj ObjUsecase
  obj.UsecaseName = vo.Naming(usecaseName)

  return &obj, nil
}

func (o ObjUsecase) GetData(PackagePath string) *ObjDataUsecase {
  return &ObjDataUsecase{
    PackagePath: PackagePath,
    UsecaseName: o.UsecaseName.String(),
  }
}

func GetUsecaseRootFolderName(o ObjUsecase) string {
  return fmt.Sprintf("usecase/%s", o.UsecaseName.LowerCase())
}

func GetInportFileName(o ObjUsecase) string {
  return fmt.Sprintf("%s/inport.go", GetUsecaseRootFolderName(o))
}

func GetOutportFileName(o ObjUsecase) string {
  return fmt.Sprintf("%s/outport.go", GetUsecaseRootFolderName(o))
}

func GetInteractorFileName(o ObjUsecase) string {
  return fmt.Sprintf("%s/interactor.go", GetUsecaseRootFolderName(o))
}
