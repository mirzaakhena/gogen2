package prod

import (
	"context"
	"github.com/mirzaakhena/gogen2/infrastructure/templates"
)

func (r *prodGateway) GetInportTemplate(ctx context.Context) string {
	return templates.InportFile
}

func (r *prodGateway) GetOutportTemplate(ctx context.Context) string {
	return templates.OutportFile
}

func (r *prodGateway) GetInteractorTemplate(ctx context.Context) string {
	return templates.InteractorFile
}
