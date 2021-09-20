package prod

import (
	"context"
	"github.com/mirzaakhena/gogen2/infrastructure/templates"
)

// GetEntityTemplate ...
func (r *prodGateway) GetEntityTemplate(ctx context.Context) string {
	return templates.EntityFile
}
