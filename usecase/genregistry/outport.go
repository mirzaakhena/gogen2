package genregistry

import (
  "context"
  "github.com/mirzaakhena/gogen2/domain/entity"
)

// Outport of GenRegistry
type Outport interface {
  FindObjController(ctx context.Context, controllerName string) (*entity.ObjController, error)
  FindAllObjGateway(ctx context.Context) ([]*entity.ObjGateway, error)
  FindObjGateway(ctx context.Context, gatewayName string) (*entity.ObjGateway, error)
  FindAllObjUsecases(ctx context.Context, objController *entity.ObjController) ([]*entity.ObjUsecase, error)
}
