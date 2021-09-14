package genrepository

import (
  "context"
  "fmt"
  "github.com/mirzaakhena/gogen2/domain/entity"
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

  objRepoObj, err := entity.NewObjRepository(req.RepositoryName, req.EntityName, req.UsecaseName, packagePath)
  if err != nil {
    return nil, err
  }

  // create folder repository
  rootFolderName := objRepoObj.GetRootFolderName()
  err = r.outport.CreateFolderIfNotExist(ctx, rootFolderName)
  if err != nil {
    return nil, err
  }

  //TODO create entity here

  existingFile := fmt.Sprintf("domain/repository/repository.go")

  // create repository.go if not exist yet
  if !r.outport.IsFileExist(ctx, existingFile) {
    repoTemplateFile := r.outport.GetRepositoryTemplateFile(ctx)
    err = r.outport.WriteFile(ctx, repoTemplateFile, existingFile, objRepoObj)
    if err != nil {
      return nil, err
    }

  } else

  {
    exist, err := objRepoObj.IsRepoExist()
    if err != nil {
      return nil, err
    }

    if exist {
      return nil, fmt.Errorf("repo is already exist")
    }

  }


  // check the prefix and give specific template for it
  constTemplateCode, err := r.outport.PrepareRepoTemplate(ctx)
  if err != nil {
    return nil, err
  }

  err = objRepoObj.InjectCode(constTemplateCode)
  if err != nil {
    return nil, err
  }

  return res, nil
}
