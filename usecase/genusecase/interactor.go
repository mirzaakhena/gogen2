package genusecase

import (
	"context"
	"fmt"
	"github.com/mirzaakhena/gogen2/domain/entity"
)

//go:generate mockery --name Outport -output mocks/

type genUsecaseInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase GenUsecase
func NewUsecase(outputPort Outport) Inport {
	return &genUsecaseInteractor{
		outport: outputPort,
	}
}

// Execute the usecase GenUsecase
func (r *genUsecaseInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	// buat object usecase
	obj, err := entity.NewObjUsecase(req.UsecaseName)
	if err != nil {
		return nil, err
	}

	// create folder usecase
	rootFolderName := entity.GetUsecaseRootFolderName(*obj)
	{
		exist, err := r.outport.CreateFolderIfNotExist(ctx, rootFolderName)
		if err != nil {
			return nil, err
		}

		if exist {
			res.Message = fmt.Sprintf("Usecase with package name %s already exist", obj.UsecaseName.LowerCase())
			return res, nil
		}
	}

	packagePath := r.outport.GetPackagePath(ctx)

	// create file inport.go
	{
		inportTemplateFile := r.outport.GetInportTemplate(ctx)
		outputFile := entity.GetInportFileName(*obj)
		_, err := r.outport.WriteFileIfNotExist(ctx, inportTemplateFile, outputFile, obj.GetData(packagePath))
		if err != nil {
			return nil, err
		}
	}

	// create file outport.go
	{
		outportTemplateFile := r.outport.GetOutportTemplate(ctx)
		outputFile := entity.GetOutportFileName(*obj)
		_, err := r.outport.WriteFileIfNotExist(ctx, outportTemplateFile, outputFile, obj.GetData(packagePath))
		if err != nil {
			return nil, err
		}
	}

	// create file interactor.go
	{
		interactorTemplateFile := r.outport.GetInteractorTemplate(ctx)
		outputFile := entity.GetInteractorFileName(*obj)
		_, err := r.outport.WriteFileIfNotExist(ctx, interactorTemplateFile, outputFile, obj.GetData(packagePath))
		if err != nil {
			return nil, err
		}

		// reformat interactor.go
		err = r.outport.Reformat(ctx, outputFile, nil)
		if err != nil {
			return nil, err
		}
	}

	return res, nil
}
