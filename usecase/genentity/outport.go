package genentity

import (
	"context"
	"github.com/mirzaakhena/gogen2/domain/service"
)

// Outport of GenEntity
type Outport interface {
	service.CreateFolderIfNotExistService
	service.WriteFileIfNotExistService
	service.ReformatService

	GetEntityTemplate(ctx context.Context) string
}
