package gentest

import (
	"context"
	"github.com/mirzaakhena/gogen2/domain/service"
)

// Outport of GenTest
type Outport interface {
	service.LogActionInterface
	service.IsFileExistService
	service.WriteFileService
	service.ReformatService
	service.GetPackagePathService

	GetTestTemplate(ctx context.Context) string
}
