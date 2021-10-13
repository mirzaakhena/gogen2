package entity

import (
  "fmt"
  "github.com/mirzaakhena/gogen2/domain/vo"
)

type ObjRegistry struct {
  RegistryName  vo.Naming
  ObjController *ObjController
  ObjGateway    *ObjGateway
  UsecaseNames  []string
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
  UsecaseNames  []string
}

func NewObjRegistry(req ObjGatewayRequest) (*ObjRegistry, error) {

  if req.ObjController == nil {
    return nil, fmt.Errorf("ObjController must not be nil")
  }

  if req.ObjGateway == nil {
    return nil, fmt.Errorf("ObjGateway must not be nil")
  }

  if len(req.UsecaseNames) == 0 {
    return nil, fmt.Errorf("usecases must not empty")
  }

  var obj ObjRegistry
  obj.RegistryName = vo.Naming(req.RegistryName)
  obj.ObjController = req.ObjController
  obj.ObjGateway = req.ObjGateway
  obj.UsecaseNames = req.UsecaseNames

  return &obj, nil
}

// GetData ...
func (o ObjRegistry) GetData(PackagePath string) *ObjDataRegistry {

  return &ObjDataRegistry{
    PackagePath:    PackagePath,
    RegistryName:   o.RegistryName.String(),
    ControllerName: o.ObjController.ControllerName.String(),
    GatewayName:    o.ObjGateway.GatewayName.LowerCase(),
    UsecaseNames:   o.UsecaseNames,
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
