package commandline

import (
	"context"
	"fmt"
	"github.com/mirzaakhena/gogen2/infrastructure/log"
	"github.com/mirzaakhena/gogen2/usecase/gencontroller"
)

// genControllerHandler ...
func (r *Controller) genControllerHandler(inputPort gencontroller.Inport) func(...string) error {

	return func(commands ...string) error {

		ctx := log.Context(context.Background())

		if len(commands) < 2 {
			err := fmt.Errorf("invalid gogen controller command format. Try this `gogen controller ControllerName UsecaseName`")
			return err
		}

		var req gencontroller.InportRequest
		req.ControllerName = commands[0]
		req.UsecaseName = commands[1]

		_, err := inputPort.Execute(ctx, req)
		if err != nil {
			return err
		}

		return nil

	}
}
