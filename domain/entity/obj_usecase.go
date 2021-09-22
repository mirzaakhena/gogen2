package entity

import (
	"fmt"
	"github.com/mirzaakhena/gogen2/application/apperror"
	"github.com/mirzaakhena/gogen2/domain/vo"
)

const (
	// OutportInterfaceName ...
	OutportInterfaceName = "Outport"
)

// ObjUsecase ...
type ObjUsecase struct {
	UsecaseName vo.Naming
}

// ObjDataUsecase ...
type ObjDataUsecase struct {
	PackagePath string
	UsecaseName string
}

// NewObjUsecase ...
func NewObjUsecase(usecaseName string) (*ObjUsecase, error) {

	if usecaseName == "" {
		return nil, apperror.UsecaseNameMustNotEmpty
	}

	var obj ObjUsecase
	obj.UsecaseName = vo.Naming(usecaseName)

	return &obj, nil
}

// GetData ...
func (o ObjUsecase) GetData(PackagePath string) *ObjDataUsecase {
	return &ObjDataUsecase{
		PackagePath: PackagePath,
		UsecaseName: o.UsecaseName.String(),
	}
}

// TODO create new usecase will create new interactor

// GetUsecaseRootFolderName ...
func GetUsecaseRootFolderName(o ObjUsecase) string {
	return fmt.Sprintf("usecase/%s", o.UsecaseName.LowerCase())
}

// GetInportFileName ...
func GetInportFileName(o ObjUsecase) string {
	return fmt.Sprintf("%s/inport.go", GetUsecaseRootFolderName(o))
}

// GetOutportFileName ...
func GetOutportFileName(o ObjUsecase) string {
	return fmt.Sprintf("%s/outport.go", GetUsecaseRootFolderName(o))
}

// GetInteractorFileName ...
func GetInteractorFileName(o ObjUsecase) string {
	return fmt.Sprintf("%s/interactor.go", GetUsecaseRootFolderName(o))
}
