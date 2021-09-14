package gentest

import (
  "context"
  "github.com/mirzaakhena/gogen2/domain/service"
)

// Outport of GenTest
type Outport interface {
  service.CreateFolderIfNotExistService
  service.WriteFileIfNotExistService
  service.IsFileExistService
  service.WriteFileService
  service.ReformatService
  service.GetPackagePathService

  GetTestTemplateFile(ctx context.Context) string


}
