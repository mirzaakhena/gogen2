package genrepository

import (
	"context"
	"fmt"
	"github.com/mirzaakhena/gogen2/domain/entity"
	"github.com/mirzaakhena/gogen2/usecase/genentity"
)

//go:generate mockery --name Outport -output mocks/

type genRepositoryInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase GenRepository
func NewUsecase(outputPort Outport) Inport {
	return &genRepositoryInteractor{
		outport: outputPort,
	}
}

// Execute the usecase GenRepository
func (r *genRepositoryInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	packagePath := r.outport.GetPackagePath(ctx)

	obj, err := entity.NewObjRepository(req.RepositoryName, req.EntityName, req.UsecaseName, packagePath)
	if err != nil {
		return nil, err
	}

	// create folder repository
	rootFolderName := obj.GetRootFolderName()
	{
		_, err := r.outport.CreateFolderIfNotExist(ctx, rootFolderName)
		if err != nil {
			return nil, err
		}
	}

	// create entity
	{
		genentity.NewUsecase(r.outport).Execute(ctx, genentity.InportRequest{EntityName: req.EntityName})
	}

	existingFile := obj.GetRepositoryFileName() // fmt.Sprintf("domain/repository/repository.go")

	// create repository.go if not exist yet
	if !r.outport.IsFileExist(ctx, existingFile) {
		repoTemplateFile := r.outport.GetRepositoryTemplate(ctx)
		err = r.outport.WriteFile(ctx, repoTemplateFile, existingFile, obj)
		if err != nil {
			return nil, err
		}

	} else {
		exist, err := obj.IsRepoExist()
		if err != nil {
			return nil, err
		}

		if exist {
			return nil, fmt.Errorf("repo is already exist")
		}

	}

	// check the prefix and give specific template for it
	constTemplateCode, err := r.outport.GetRepositoryFunctionTemplate(ctx)
	if err != nil {
		return nil, err
	}

	err = obj.InjectCode(constTemplateCode)
	if err != nil {
		return nil, err
	}

	// reformat interactor.go
	err = r.outport.Reformat(ctx, existingFile)
	if err != nil {
		return nil, err
	}

	return res, nil
}
