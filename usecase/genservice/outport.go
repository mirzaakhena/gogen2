package genservice

import (
  "context"
  "github.com/mirzaakhena/gogen2/domain/service"
)

// Outport of GenService
type Outport interface {
  service.PrintTemplateService
  service.ReformatService
  service.GetPackagePathService

  GetServiceInterfaceTemplate(ctx context.Context) string
  GetServiceInjectTemplate(ctx context.Context) string
}
