package genentity

import (
	"context"
	"github.com/mirzaakhena/gogen2/domain/entity"
)

//go:generate mockery --name Outport -output mocks/

type genEntityInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase GenEntity
func NewUsecase(outputPort Outport) Inport {
	return &genEntityInteractor{
		outport: outputPort,
	}
}

// Execute the usecase GenEntity
func (r *genEntityInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	// buat object entity
	obj, err := entity.NewObjEntity(req.EntityName)
	if err != nil {
		return nil, err
	}

	// create folder entity
	rootFolderName := obj.GetRootFolderName()
	{
		_, err := r.outport.CreateFolderIfNotExist(ctx, rootFolderName)
		if err != nil {
			return nil, err
		}
	}

	exist, err := obj.IsExist()
	if err != nil {
		return nil, err
	}

	// entity already exist, nothing to do
	if exist {
		res.Message = "Entity already exist"
		return nil, nil
	}

	// create file entity.go
	{
		inportTemplateFile := r.outport.GetEntityTemplate(ctx)
		outputFile := obj.GetEntityFileName()
		_, err := r.outport.WriteFileIfNotExist(ctx, inportTemplateFile, outputFile, obj.GetData())
		if err != nil {
			return nil, err
		}

		// reformat interactor.go
		err = r.outport.Reformat(ctx, outputFile)
		if err != nil {
			return nil, err
		}

	}

	return res, nil
}
