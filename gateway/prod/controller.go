package prod

import (
  "context"
  "github.com/mirzaakhena/gogen2/infrastructure/templates"
)

func (r *prodGateway) GetResponseTemplate(ctx context.Context) string {
  return templates.ControllerGinGonicResponseFile
}

func (r *prodGateway) GetInterceptorTemplate(ctx context.Context, framework string) string {
  return templates.ControllerGinGonicInterceptorFile
}

func (r *prodGateway) GetRouterTemplate(ctx context.Context, framework string) string {
  return templates.ControllerGinGonicRouterStructFile
}

func (r *prodGateway) GetHandlerTemplate(ctx context.Context, framework string) string {
  return templates.ControllerGinGonicHandlerFuncFile
}

func (r *prodGateway) GetRouterRegisterTemplate(ctx context.Context) string {
  return templates.ControllerGinGonicRouterRegisterFile
}

func (r *prodGateway) GetRouterInportTemplate(ctx context.Context) string {
  return templates.ControllerGinGonicRouterInportFile
}
