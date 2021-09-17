package genlog

import (
	"context"
	"github.com/mirzaakhena/gogen2/domain/service"
)

// Outport of GenLog
type Outport interface {
	service.CreateFolderIfNotExistService
	service.WriteFileIfNotExistService

	GetLogInterfaceTemplate(ctx context.Context) string
	GetLogImplementationFileName(ctx context.Context) string
}
