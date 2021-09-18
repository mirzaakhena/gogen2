package gentest

import (
	"context"
	"github.com/mirzaakhena/gogen2/domain/service"
	"github.com/mirzaakhena/gogen2/usecase/genlog"
)

// Outport of GenTest
type Outport interface {
	genlog.Outport
	service.IsFileExistService
	service.WriteFileService
	service.ReformatService
	service.GetPackagePathService

	GetTestTemplate(ctx context.Context) string
}