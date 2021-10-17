package registry

import (
  "flag"
  "fmt"
  "github.com/mirzaakhena/gogen2/application"
  "github.com/mirzaakhena/gogen2/controller"
  "github.com/mirzaakhena/gogen2/controller/commandline"
  "github.com/mirzaakhena/gogen2/gateway/prod"
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

type gogen2 struct {
  CommandMap map[string]func(...string) error
  controller.Controller
}

// NewGogen2 ...
func NewGogen2() func() application.RegistryContract {
  return func() application.RegistryContract {

    datasource := prod.NewProdGateway()

    commandMap := make(map[string]func(...string) error, 0)

    return &gogen2{
      CommandMap: commandMap,
      Controller: &commandline.Controller{
        CommandMap:           commandMap,
        GenUsecaseInport:     genusecase.NewUsecase(datasource),
        GenTestInport:        gentest.NewUsecase(datasource),
        GenEntityInport:      genentity.NewUsecase(datasource),
        GenRepositoryInport:  genrepository.NewUsecase(datasource),
        GenServiceInport:     genservice.NewUsecase(datasource),
        GenGatewayInport:     gengateway.NewUsecase(datasource),
        GenErrorInport:       generror.NewUsecase(datasource),
        GenControllerInport:  gencontroller.NewUsecase(datasource),
        GenRegistryInport:    genregistry.NewUsecase(datasource),
        GenValueObjectInport: genvalueobject.NewUsecase(datasource),
        GenValueStringInport: genvaluestring.NewUsecase(datasource),
      },
    }

  }
}

// RunApplication ...
func (r *gogen2) RunApplication() {
  flag.Parse()
  cmd := flag.Arg(0)

  if cmd == "" {
    fmt.Printf("try gogen2 usecase\n")
    return
  }

  var values = make([]string, 0)
  if flag.NArg() > 1 {
    values = flag.Args()[1:]
  }

  f, exists := r.CommandMap[cmd]
  if !exists {
    fmt.Printf("Command %s is not recognized\n", cmd)
    return
  }
  err := f(values...)
  if err != nil {
    fmt.Printf("%s\n", err.Error())
    return
  }
}
