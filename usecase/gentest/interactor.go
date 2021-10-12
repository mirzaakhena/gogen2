package gentest

import (
	"context"
	"fmt"
	"github.com/mirzaakhena/gogen2/domain/entity"
	"github.com/mirzaakhena/gogen2/domain/service"
	"github.com/mirzaakhena/gogen2/domain/vo"
)

//go:generate mockery --name Outport -output mocks/

type genTestInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase GenTest
func NewUsecase(outputPort Outport) Inport {
	return &genTestInteractor{
		outport: outputPort,
	}
}

// Execute the usecase GenTest
func (r *genTestInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	// create log
	{
		err := service.ConstructLog(ctx, r.outport)
		if err != nil {
			return nil, err
		}
	}

	objUsecase, err := entity.NewObjUsecase(req.UsecaseName)
	if err != nil {
		return nil, err
	}

	obj, err := entity.NewObjTesting(req.TestName, *objUsecase)
	if err != nil {
		return nil, err
	}

	// create interactor_test.go

	outputFile := obj.GetTestFileName()

	if r.outport.IsFileExist(ctx, outputFile) {
		res.Message = fmt.Sprintf("file test %s already exists", req.TestName)
		return res, nil
	}

	packagePath := r.outport.GetPackagePath(ctx)

	outportMethods, err := vo.NewOutportMethods(req.UsecaseName, packagePath)
	if err != nil {
		return nil, err
	}

	tmp := r.outport.GetTestTemplate(ctx)

	err = r.outport.WriteFile(ctx, tmp, outputFile, obj.GetData(packagePath, outportMethods))
	if err != nil {
		return nil, err
	}

	// reformat interactor.go
	err = r.outport.Reformat(ctx, outputFile, nil)
	if err != nil {
		return nil, err
	}

	return res, nil
}
