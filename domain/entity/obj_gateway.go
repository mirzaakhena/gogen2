package entity

import (
	"fmt"
	"github.com/mirzaakhena/gogen2/application/apperror"
	"github.com/mirzaakhena/gogen2/domain/vo"
)

// ObjGateway  depend on (which) usecase that want to be tested
type ObjGateway struct {
	GatewayName vo.Naming
	ObjUsecase  ObjUsecase
}

// ObjDataGateway  ...
type ObjDataGateway struct {
	PackagePath string
	UsecaseName string
	GatewayName string
	Methods     vo.OutportMethods
}

// NewObjGateway   ...
func NewObjGateway(gatewayName string, objUsecase ObjUsecase) (*ObjGateway, error) {

	if gatewayName == "" {
		return nil, apperror.GatewayNameMustNotEmpty
	}

	var obj ObjGateway
	obj.GatewayName = vo.Naming(gatewayName)
	obj.ObjUsecase = objUsecase

	return &obj, nil
}

// GetData ...
func (o ObjGateway) GetData(PackagePath string, outportMethods vo.OutportMethods) *ObjDataGateway {
	return &ObjDataGateway{
		PackagePath: PackagePath,
		UsecaseName: o.ObjUsecase.UsecaseName.String(),
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
