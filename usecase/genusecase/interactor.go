package genusecase

import (
  "context"
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

  packagePath := r.outport.GetPackagePath(ctx)

  // buat object usecase
  obj, err := entity.NewObjUsecase(req.UsecaseName, packagePath)
  if err != nil {
    return nil, err
  }

  // create folder usecase
  rootFolderName := obj.GetRootFolderName()
  err = r.outport.CreateFolderIfNotExist(ctx, rootFolderName)
  if err != nil {
    return nil, err
  }

  // create file inport.go
  {
    inportTemplateFile := r.outport.GetInportTemplateFile(ctx)
    outputFile := obj.GetInportFileName()
    err := r.outport.WriteFileIfNotExist(ctx, inportTemplateFile, outputFile, obj.GetData())
    if err != nil {
      return nil, err
    }
  }

  // create file outport.go
  {
    outportTemplateFile := r.outport.GetOutportTemplateFile(ctx)
    outputFile := obj.GetOutportFileName()
    err := r.outport.WriteFileIfNotExist(ctx, outportTemplateFile, outputFile, obj.GetData())
    if err != nil {
      return nil, err
    }
  }

  // create file interactor.go
  {
    interactorTemplateFile := r.outport.GetInteractorTemplateFile(ctx)
    outputFile := obj.GetInteractorFileName()
    err := r.outport.WriteFileIfNotExist(ctx, interactorTemplateFile, outputFile, obj.GetData())
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
