package genusecase

import (
  "context"
  "github.com/mirzaakhena/gogen2/domain/service"
)

// Outport of GenUsecase
type Outport interface {
  service.CreateFolderIfNotExistService
  service.WriteFileIfNotExistService
  service.ReformatService
  service.GetPackagePathService

  GetInportTemplateFile(ctx context.Context) string
  GetOutportTemplateFile(ctx context.Context) string
  GetInteractorTemplateFile(ctx context.Context) string
}
