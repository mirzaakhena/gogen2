package gencontroller

import (
  "context"
  "github.com/mirzaakhena/gogen2/domain/service"
)

// Outport of GenController
type Outport interface {
  service.ApplicationActionInterface
  service.LogActionInterface

  service.IsFileExistService
  service.WriteFileService
  service.ReformatService
  service.GetPackagePathService
  service.PrintTemplateService

  GetResponseTemplate(ctx context.Context) string
  GetInterceptorTemplate(ctx context.Context, framework string) string
  GetRouterTemplate(ctx context.Context, framework string) string
  GetHandlerTemplate(ctx context.Context, framework string) string
  GetInportTemplate(ctx context.Context) string
  GetRouterRegisterTemplate(ctx context.Context) string
}
