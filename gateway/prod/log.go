package prod

import (
	"context"
	"github.com/mirzaakhena/gogen2/infrastructure/templates"
)

// GetLogInterfaceTemplate ...
func (r *prodGateway) GetLogInterfaceTemplate(ctx context.Context) string {
	return templates.LogFile
}

// GetLogImplementationTemplate ...
func (r *prodGateway) GetLogImplementationTemplate(ctx context.Context) string {
	return templates.LogDefaultFile
}
