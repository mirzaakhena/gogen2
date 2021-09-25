package genregistry

import (
  "context"
  "fmt"
  "github.com/mirzaakhena/gogen2/domain/entity"
)

//go:generate mockery --name Outport -output mocks/

type genRegistryInteractor struct {
  outport Outport
}

// NewUsecase is constructor for create default implementation of usecase GenRegistry
func NewUsecase(outputPort Outport) Inport {
  return &genRegistryInteractor{
    outport: outputPort,
  }
}

// Execute the usecase GenRegistry
func (r *genRegistryInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

  res := &InportResponse{}

  objController, err := r.outport.FindObjController(ctx, req.ControllerName)
  if err != nil {
    return nil, err
  }

  var objGateway *entity.ObjGateway

  if req.GatewayName == "" {

    objGateways, err := r.outport.FindAllObjGateway(ctx)
    if err != nil {
      return nil, err
    }

    if len(objGateways) > 1 {
      names := make([]string, 0)
      for _, g := range objGateways {
        names = append(names, g.GatewayName.String())
      }
      return nil, fmt.Errorf("select one of this gateways %v", names)
    }

    if len(objGateways) == 0 {
      return nil, fmt.Errorf("no gateway found")
    }

    objGateway = objGateways[0]

  } else {
    objGateway, err = r.outport.FindObjGateway(ctx, req.GatewayName)
    if err != nil {
      return nil, err
    }

    if objGateway == nil {
      return nil, fmt.Errorf("no gateway with name %s found", req.GatewayName)
    }
  }

  objRegistry, err := entity.NewObjRegistry(ctx, entity.ObjGatewayRequest{
    RegistryName:   req.RegistryName,
    ObjController:  objController,
    ObjGateway:     objGateway,
    FindAllUsecase: r.outport,
  })
  if err != nil {
    return nil, err
  }

  _ = objRegistry

  return res, nil
}
