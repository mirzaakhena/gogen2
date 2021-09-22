package generror

import (
	"context"
	"github.com/mirzaakhena/gogen2/domain/service"
)

// Outport of GenError
type Outport interface {
	service.PrintTemplateService
	service.ApplicationActionInterface
	service.ReformatService

	GetErrorLineTemplate(ctx context.Context) string
}
