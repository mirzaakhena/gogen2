package commandline

import (
  "github.com/mirzaakhena/gogen2/usecase/gentest"
  "github.com/mirzaakhena/gogen2/usecase/genusecase"
)

type Controller struct {
  CommandMap       map[string]func(...string) error
  GenUsecaseInport genusecase.Inport
  GenTestInport    gentest.Inport
}

// RegisterRouter registering all the router
func (r *Controller) RegisterRouter() {
  r.CommandMap["usecase"] = r.genUsecaseHandler(r.GenUsecaseInport)
  r.CommandMap["test"] = r.genTestHandler(r.GenTestInport)
}
