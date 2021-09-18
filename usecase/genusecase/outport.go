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

  GetInportTemplate(ctx context.Context) string
  GetOutportTemplate(ctx context.Context) string
  GetInteractorTemplate(ctx context.Context) string
}
