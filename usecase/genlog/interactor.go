package genlog

import (
	"context"
	"github.com/mirzaakhena/gogen2/domain/entity"
)

//go:generate mockery --name Outport -output mocks/

type genLogInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase GenLog
func NewUsecase(outputPort Outport) Inport {
	return &genLogInteractor{
		outport: outputPort,
	}
}

// Execute the usecase GenLog
func (r *genLogInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	obj, err := entity.NewObjLog()
	if err != nil {
		return nil, err
	}

	rootFolderName := obj.GetRootFolderName()
	{
		_, err := r.outport.CreateFolderIfNotExist(ctx, rootFolderName)
		if err != nil {
			return nil, err
		}
	}

	{
		logTemplateFile := r.outport.GetLogInterfaceTemplate(ctx)
		outputFile := obj.GetLogInterfaceFileName()
		_, err := r.outport.WriteFileIfNotExist(ctx, logTemplateFile, outputFile, struct{}{})
		if err != nil {
			return nil, err
		}
	}

	{
		logTemplateFile := r.outport.GetLogImplementationFileName(ctx)
		outputFile := obj.GetLogImplementationFileName()
		_, err := r.outport.WriteFileIfNotExist(ctx, logTemplateFile, outputFile, struct{}{})
		if err != nil {
			return nil, err
		}
	}

	return res, nil
}
