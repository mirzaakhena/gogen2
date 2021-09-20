package prod

import (
	"context"
	"github.com/mirzaakhena/gogen2/infrastructure/templates"
)

// GetInportTemplate ...
func (r *prodGateway) GetInportTemplate(ctx context.Context) string {
	return templates.InportFile
}

// GetOutportTemplate ...
func (r *prodGateway) GetOutportTemplate(ctx context.Context) string {
	return templates.OutportFile
}

// GetInteractorTemplate ...
func (r *prodGateway) GetInteractorTemplate(ctx context.Context) string {
	return templates.InteractorFile
}
