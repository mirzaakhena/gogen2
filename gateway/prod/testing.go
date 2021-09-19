package prod

import (
	"context"
	"github.com/mirzaakhena/gogen2/infrastructure/templates"
)

func (r *prodGateway) GetTestTemplate(ctx context.Context) string {
	return templates.TestFile
}
