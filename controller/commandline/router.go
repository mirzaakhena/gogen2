package commandline

import (
  "github.com/mirzaakhena/gogen2/usecase/gencontroller"
  "github.com/mirzaakhena/gogen2/usecase/genentity"
  "github.com/mirzaakhena/gogen2/usecase/generror"
  "github.com/mirzaakhena/gogen2/usecase/gengateway"
  "github.com/mirzaakhena/gogen2/usecase/genregistry"
  "github.com/mirzaakhena/gogen2/usecase/genrepository"
  "github.com/mirzaakhena/gogen2/usecase/genservice"
  "github.com/mirzaakhena/gogen2/usecase/gentest"
  "github.com/mirzaakhena/gogen2/usecase/genusecase"
  "github.com/mirzaakhena/gogen2/usecase/genvalueobject"
  "github.com/mirzaakhena/gogen2/usecase/genvaluestring"
)

// Controller ...
type Controller struct {
  CommandMap           map[string]func(...string) error
  GenUsecaseInport     genusecase.Inport
  GenTestInport        gentest.Inport
  GenEntityInport      genentity.Inport
  GenRepositoryInport  genrepository.Inport
  GenServiceInport     genservice.Inport
  GenGatewayInport     gengateway.Inport
  GenErrorInport       generror.Inport
  GenControllerInport  gencontroller.Inport
  GenRegistryInport    genregistry.Inport
  GenValueObjectInport genvalueobject.Inport
  GenValueStringInport genvaluestring.Inport
}

// RegisterRouter registering all the router
func (r *Controller) RegisterRouter() {
  r.CommandMap["usecase"] = r.genUsecaseHandler(r.GenUsecaseInport)
  r.CommandMap["test"] = r.genTestHandler(r.GenTestInport)
  r.CommandMap["entity"] = r.genEntityHandler(r.GenEntityInport)
  r.CommandMap["repository"] = r.genRepositoryHandler(r.GenRepositoryInport)
  r.CommandMap["service"] = r.genServiceHandler(r.GenServiceInport)
  r.CommandMap["gateway"] = r.genGatewayHandler(r.GenGatewayInport)
  r.CommandMap["error"] = r.genErrorHandler(r.GenErrorInport)
  r.CommandMap["controller"] = r.genControllerHandler(r.GenControllerInport)
  r.CommandMap["registry"] = r.genRegistryHandler(r.GenRegistryInport)
  r.CommandMap["valueobject"] = r.genValueObjectHandler(r.GenValueObjectInport)
  r.CommandMap["valuestring"] = r.genValueStringHandler(r.GenValueStringInport)
}
