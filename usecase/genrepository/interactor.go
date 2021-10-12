package genrepository

import (
	"context"
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

	obj, err := entity.NewObjRepository(req.RepositoryName, req.EntityName, req.UsecaseName)
	if err != nil {
		return nil, err
	}

	// create folder repository
	{
		_, err := r.outport.CreateFolderIfNotExist(ctx, "domain/repository")
		if err != nil {
			return nil, err
		}
	}

	// create entity
	{
		_, err := genentity.NewUsecase(r.outport).Execute(ctx, genentity.InportRequest{EntityName: req.EntityName})
		if err != nil {
			return nil, err
		}
	}

	existingFile := "domain/repository/repository.go"

	// create repository.go file if not exist yet
	if !r.outport.IsFileExist(ctx, existingFile) {
		tmp := r.outport.GetRepositoryTemplate(ctx)
		err := r.outport.WriteFile(ctx, tmp, existingFile, struct {}{})
		if err != nil {
			return nil, err
		}
	}

	// repository.go file is already exist, but is the innterface is exist ?
	exist, err := obj.IsRepoExist()
	if err != nil {
		return nil, err
	}

	packagePath := r.outport.GetPackagePath(ctx)

	if !exist {
		// check the prefix and give specific template for it
		templateCode, err := r.outport.GetRepositoryFunctionTemplate(ctx, obj.RepositoryName)
		if err != nil {
			return nil, err
		}

		templateHasBeenInjected, err := r.outport.PrintTemplate(ctx, templateCode, obj.GetData(packagePath))
		if err != nil {
			return nil, err
		}

		bytes, err := obj.InjectCode(templateHasBeenInjected)
		if err != nil {
			return nil, err
		}

		// reformat interactor.go
		err = r.outport.Reformat(ctx, existingFile, bytes)
		if err != nil {
			return nil, err
		}
	}

	// if usecase name is not empty means we want to inject to usecase
	if obj.ObjUsecase.UsecaseName.IsEmpty() {
		return res, nil
	}

	{
		// inject to outport
		err := obj.InjectToOutport()
		if err != nil {
			return nil, err
		}

		outportFile := obj.ObjUsecase.GetOutportFileName()

		// reformat outport.go
		err = r.outport.Reformat(ctx, outportFile, nil)
		if err != nil {
			return nil, err
		}
	}


	{
		// check the prefix and give specific template for it
		interactorCode, err := r.outport.GetInteractorRepoCallTemplate(ctx, obj.RepositoryName)
		if err != nil {
			return nil, err
		}

		templateHasBeenInjected, err := r.outport.PrintTemplate(ctx, interactorCode, obj.GetData(packagePath))
		if err != nil {
			return nil, err
		}

		interactorBytes, err := obj.InjectToInteractor(templateHasBeenInjected)
		if err != nil {
			return nil, err
		}

		// reformat interactor.go
		err = r.outport.Reformat(ctx, obj.ObjUsecase.GetInteractorFileName(), interactorBytes)
		if err != nil {
			return nil, err
		}
	}


	return res, nil
}
