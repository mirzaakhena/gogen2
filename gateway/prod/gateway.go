package prod

import (
	"context"
	"github.com/mirzaakhena/gogen2/infrastructure/templates"
)

// GetGatewayTemplate ...
func (r *prodGateway) GetGatewayTemplate(ctx context.Context) string {
	return templates.GatewayGormFile
}

// GetGatewayMethodTemplate ...
func (r *prodGateway) GetGatewayMethodTemplate(ctx context.Context) string {
	return templates.GatewayGormInjectMethodFile
}
