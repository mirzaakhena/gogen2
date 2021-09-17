package gentest

import (
  "context"
  "github.com/mirzaakhena/gogen2/domain/entity"
  "github.com/mirzaakhena/gogen2/usecase/genlog"
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

  packagePath := r.outport.GetPackagePath(ctx)

  // create log
  {
    _, err := genlog.NewUsecase(r.outport).Execute(ctx, genlog.InportRequest{})
    if err != nil {
      return nil, err
    }
  }

  obj, err := entity.NewObjTesting(req.TestName, req.UsecaseName, packagePath)
  if err != nil {
    return nil, err
  }

  // create interactor_test.go
  {
    outputFile := obj.GetTestFileName()

    if !r.outport.IsFileExist(ctx, outputFile) {

      testTemplateFile := r.outport.GetTestTemplateFile(ctx)
      err := r.outport.WriteFile(ctx, testTemplateFile, outputFile, obj.GetData())
      if err != nil {
        return nil, err
      }

      // reformat interactor.go
      err = r.outport.Reformat(ctx, outputFile)
      if err != nil {
        return nil, err
      }

    }

  }

  return res, nil
}
