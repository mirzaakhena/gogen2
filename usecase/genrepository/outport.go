package genrepository

import (
  "context"
  "github.com/mirzaakhena/gogen2/domain/service"
  "github.com/mirzaakhena/gogen2/domain/vo"
  "github.com/mirzaakhena/gogen2/usecase/genentity"
)

// Outport of GenRepository
type Outport interface {
  genentity.Outport
  service.IsFileExistService
  service.WriteFileService
  service.GetPackagePathService
  service.CreateFolderIfNotExistService
  service.PrintTemplateService

  GetRepoInterfaceTemplate(ctx context.Context, repoName vo.Naming) string
  GetRepoInjectTemplate(ctx context.Context, repoName vo.Naming) string
}
