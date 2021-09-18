package genrepository

import (
  "context"
  "github.com/mirzaakhena/gogen2/domain/service"
  "github.com/mirzaakhena/gogen2/usecase/genentity"
)

// Outport of GenRepository
type Outport interface {
  genentity.Outport
  service.IsFileExistService
  service.WriteFileService
  service.GetPackagePathService
  service.CreateFolderIfNotExistService
  GetRepositoryTemplate(ctx context.Context) string
  GetRepositoryFunctionTemplate(ctx context.Context) (string, error)
}

