package registry

import (
	"context"
	"flag"
	"github.com/mirzaakhena/gogen2/application"
	"github.com/mirzaakhena/gogen2/controller/commandline"
	"github.com/mirzaakhena/gogen2/gateway/prod"
	"github.com/mirzaakhena/gogen2/infrastructure/log"
	"github.com/mirzaakhena/gogen2/usecase/genentity"
	"github.com/mirzaakhena/gogen2/usecase/generror"
	"github.com/mirzaakhena/gogen2/usecase/gengateway"
	"github.com/mirzaakhena/gogen2/usecase/genrepository"
	"github.com/mirzaakhena/gogen2/usecase/gentest"
	"github.com/mirzaakhena/gogen2/usecase/genusecase"
)

type gogen2 struct {
	commandlineController commandline.Controller
}

// NewGogen2 ...
func NewGogen2() func() application.RegistryContract {
	return func() application.RegistryContract {

		datasource := prod.NewProdGateway()

		return &gogen2{
			commandlineController: commandline.Controller{
				CommandMap: make(map[string]func(...string) error, 0),

				GenUsecaseInport:    genusecase.NewUsecase(datasource),
				GenTestInport:       gentest.NewUsecase(datasource),
				GenEntityInport:     genentity.NewUsecase(datasource),
				GenRepositoryInport: genrepository.NewUsecase(datasource),
				GenGatewayInport:    gengateway.NewUsecase(datasource),
				GenErrorInport:      generror.NewUsecase(datasource),
			},
		}

	}
}

// RunApplication ...
func (r *gogen2) RunApplication() {
	flag.Parse()
	cmd := flag.Arg(0)

	if cmd == "" {
		log.Error(context.Background(), "try gogen2 usecase CreateOrder")
		return
	}

	var values = make([]string, 0)
	if flag.NArg() > 1 {
		values = flag.Args()[1:]
	}

	f, exists := r.commandlineController.CommandMap[cmd]
	if !exists {
		log.Error(context.Background(), "Command %s is not recognized", cmd)
		return
	}
	err := f(values...)
	if err != nil {
		log.Error(context.Background(), err.Error())
		return
	}
}

// SetupController ...
func (r *gogen2) SetupController() {
	r.commandlineController.RegisterRouter()
}
