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

  // create object usecase
  obj, err := entity.NewObjUsecase(req.UsecaseName)
  if err != nil {
    return nil, err
  }

  // create folder usecase
  {
    folderPath := obj.GetUsecaseRootFolderName()
    _, err := r.outport.CreateFolderIfNotExist(ctx, folderPath)
    if err != nil {
      return nil, err
    }
  }

  packagePath := r.outport.GetPackagePath(ctx)

  // create file inport.go
  {
    tmp := r.outport.GetInportTemplate(ctx)
		outputFile := obj.GetInportFileName()
    _, err = r.outport.WriteFileIfNotExist(ctx, tmp, outputFile, obj.GetData(packagePath))
    if err != nil {
      return nil, err
    }
  }

  // create file outport.go
  {
		tmp := r.outport.GetOutportTemplate(ctx)
		outputFile := obj.GetOutportFileName()
		_, err = r.outport.WriteFileIfNotExist(ctx, tmp, outputFile, obj.GetData(packagePath))
		if err != nil {
			return nil, err
		}
  }

  // create file interactor.go
  {
		tmp := r.outport.GetInteractorTemplate(ctx)
		outputFile := obj.GetInteractorFileName()
		_, err = r.outport.WriteFileIfNotExist(ctx, tmp, outputFile, obj.GetData(packagePath))
		if err != nil {
			return nil, err
		}

    // reformat interactor.go since we have some import on it
    err = r.outport.Reformat(ctx, outputFile, nil)
    if err != nil {
      return nil, err
    }
  }

  return res, nil
}
