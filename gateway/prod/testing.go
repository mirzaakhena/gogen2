package prod

import (
	"context"
	"github.com/mirzaakhena/gogen2/infrastructure/templates"
)

// GetTestTemplate ...
func (r *prodGateway) GetTestTemplate(ctx context.Context) string {
	return templates.ReadFile("usecase/usecase_test._go")
}
