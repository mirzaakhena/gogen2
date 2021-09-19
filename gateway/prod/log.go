package prod

import (
  "context"
  "github.com/mirzaakhena/gogen2/infrastructure/templates"
)

func (r *prodGateway) GetLogInterfaceTemplate(ctx context.Context) string {
  return templates.LogFile
}

func (r *prodGateway) GetLogImplementationTemplate(ctx context.Context) string {
  return templates.LogDefaultFile
}
