package genrepository

import (
  "context"
  "github.com/mirzaakhena/gogen2/domain/service"
)

// Outport of GenRepository
type Outport interface {
  service.IsFileExistService
  service.WriteFileService
  service.GetPackagePathService
  service.CreateFolderIfNotExistService
  GetRepositoryTemplateFile(ctx context.Context) string
  PrepareRepoTemplate(ctx context.Context) (string, error)
}
