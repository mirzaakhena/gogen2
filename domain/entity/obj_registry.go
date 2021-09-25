package entity

import (
  "context"
  "fmt"
  "github.com/mirzaakhena/gogen2/domain/repository"
  "github.com/mirzaakhena/gogen2/domain/vo"
)

type ObjRegistry struct {
  RegistryName  vo.Naming
  ObjController *ObjController
  ObjUsecases   []*ObjUsecase
  ObjGateway    *ObjGateway
}

type ObjDataRegistry struct {
  RegistryName   string
  ControllerName string
  UsecaseNames   []string
  GatewayName    string
}

type ObjGatewayRequest struct {
  RegistryName   string
  ObjController  *ObjController
  ObjGateway     *ObjGateway
  FindAllUsecase repository.FindAllObjUsecasesRepo
}

func NewObjRegistry(ctx context.Context, req ObjGatewayRequest) (*ObjRegistry, error) {

  if req.ObjController == nil {
    return nil, fmt.Errorf("ObjController must not be nil")
  }

  var obj ObjRegistry
  obj.RegistryName = vo.Naming(req.RegistryName)
  obj.ObjController = req.ObjController
  obj.ObjGateway = req.ObjGateway
  usecases, err := req.FindAllUsecase.FindAllObjUsecases(ctx, obj.ObjController)
  if err != nil {
    return nil, err
  }

  obj.ObjUsecases = usecases

  return &obj, nil
}

