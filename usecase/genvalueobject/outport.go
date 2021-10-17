package genvalueobject

import (
  "context"
  "github.com/mirzaakhena/gogen2/domain/service"
)

// Outport of GenValueObject
type Outport interface {
  service.CreateFolderIfNotExistService
  service.WriteFileIfNotExistService

  GetValueObjectTemplate(ctx context.Context) string
  GetValueStringTemplate(ctx context.Context) string
}
