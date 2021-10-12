package commandline

import (
  "context"
  "fmt"
  "github.com/mirzaakhena/gogen2/usecase/genregistry"
)

// genRegistryHandler ...
func (r *Controller) genRegistryHandler(inputPort genregistry.Inport) func(...string) error {

  return func(commands ...string) error {

    ctx := context.Background()

    if len(commands) < 2 {
      err := fmt.Errorf("invalid gogen registry command format. Try this `gogen registry RegistryName ControllerName`")
      return err
    }

    var req genregistry.InportRequest
    req.RegistryName = commands[0]
    req.ControllerName = commands[1]

    if len(commands) >= 3 {
      req.GatewayName = commands[2]
    }

    _, err := inputPort.Execute(ctx, req)
    if err != nil {
      return err
    }

    return nil

  }
}
