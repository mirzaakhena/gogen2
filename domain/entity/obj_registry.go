package entity

import (
  "fmt"
  "github.com/mirzaakhena/gogen2/domain/vo"
)

type ObjRegistry struct {
  RegistryName  vo.Naming
  ObjController *ObjController
  ObjUsecases   []*ObjUsecase
  ObjGateway    *ObjGateway
}

type ObjDataRegistry struct {
  PackagePath    string
  RegistryName   string
  ControllerName string
  UsecaseNames   []string
  GatewayName    string
}

type ObjGatewayRequest struct {
  RegistryName  string
  ObjController *ObjController
  ObjGateway    *ObjGateway
  Usecases      []*ObjUsecase
}

func NewObjRegistry(req ObjGatewayRequest) (*ObjRegistry, error) {

  if req.ObjController == nil {
    return nil, fmt.Errorf("ObjController must not be nil")
  }

  if req.ObjGateway == nil {
    return nil, fmt.Errorf("ObjGateway must not be nil")
  }

  if req.Usecases == nil {
    return nil, fmt.Errorf("usecases must not empty")
  }

  var obj ObjRegistry
  obj.RegistryName = vo.Naming(req.RegistryName)
  obj.ObjController = req.ObjController
  obj.ObjGateway = req.ObjGateway
  obj.ObjUsecases = req.Usecases

  return &obj, nil
}

// GetData ...
func (o ObjRegistry) GetData(PackagePath string) *ObjDataRegistry {

  usecaseNames := make([]string, 0)

  for _, u := range o.ObjUsecases {
    usecaseNames = append(usecaseNames, u.UsecaseName.String())
  }

  return &ObjDataRegistry{
    PackagePath:    PackagePath,
    RegistryName:   o.RegistryName.String(),
    ControllerName: o.ObjController.ControllerName.String(),
    UsecaseNames:   usecaseNames,
  }
}

// GetRegistryRootFolderName ...
func GetRegistryRootFolderName() string {
  return fmt.Sprintf("application/registry")
}

// GetApplicationFileName ...
func GetApplicationFileName() string {
  return fmt.Sprintf("application/application.go")
}

// GetRegistryFileName ...
func GetRegistryFileName(obj ObjRegistry) string {
  return fmt.Sprintf("%s/%s.go", GetRegistryRootFolderName(), obj.RegistryName.LowerCase())
}


