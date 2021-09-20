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
	service.ApplicationActionInterface

	GetRepositoryTemplate(ctx context.Context) string
	GetRepositoryFunctionTemplate(ctx context.Context, repoName vo.Naming) (string, error)
	GetInteractorRepoCallTemplate(ctx context.Context, repoName vo.Naming) (string, error)
}
