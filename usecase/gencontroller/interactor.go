package gencontroller

import (
  "context"
  "github.com/mirzaakhena/gogen2/domain/entity"
  "github.com/mirzaakhena/gogen2/domain/service"
)

//go:generate mockery --name Outport -output mocks/

type genControllerInteractor struct {
  outport Outport
}

// NewUsecase is constructor for create default implementation of usecase GenController
func NewUsecase(outputPort Outport) Inport {
  return &genControllerInteractor{
    outport: outputPort,
  }
}

// Execute the usecase GenController
func (r *genControllerInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

  res := &InportResponse{}

  objUsecase, err := entity.NewObjUsecase(req.UsecaseName)
  if err != nil {
    return nil, err
  }

  objCtrl, err := entity.NewObjController(req.ControllerName)
  if err != nil {
    return nil, err
  }

  err = service.ConstructApplication(ctx, r.outport)
  if err != nil {
    return nil, err
  }

  err = service.ConstructLog(ctx, r.outport)
  if err != nil {
    return nil, err
  }

  _, err = r.outport.CreateFolderIfNotExist(ctx, entity.GetControllerRootFolderName(*objCtrl))
  if err != nil {
    return nil, err
  }

  packagePath := r.outport.GetPackagePath(ctx)

  framework := "gingonic"

  objDataCtrl := objCtrl.GetData(packagePath, *objUsecase)

  // response.go
  {
    filename := entity.GetControllerResponseFileName(*objCtrl)
    if !r.outport.IsFileExist(ctx, filename) {
      templateFile := r.outport.GetResponseTemplate(ctx)

      err := r.outport.WriteFile(ctx, templateFile, filename, objDataCtrl)
      if err != nil {
        return nil, err
      }

      err = r.outport.Reformat(ctx, filename, nil)
      if err != nil {
        return nil, err
      }
    }
  }

  // interceptor.go
  {
    filename := entity.GetControllerInterceptorFileName(*objCtrl)
    if !r.outport.IsFileExist(ctx, filename) {
      templateFile := r.outport.GetInterceptorTemplate(ctx, framework)

      err := r.outport.WriteFile(ctx, templateFile, filename, objDataCtrl)
      if err != nil {
        return nil, err
      }

      err = r.outport.Reformat(ctx, filename, nil)
      if err != nil {
        return nil, err
      }
    }
  }

  // handler_xxx.go
  {
    filename := entity.GetControllerHandlerFileName(*objCtrl, *objUsecase)
    if !r.outport.IsFileExist(ctx, filename) {
      templateFile := r.outport.GetHandlerTemplate(ctx, framework)

      err := r.outport.WriteFile(ctx, templateFile, filename, objDataCtrl)
      if err != nil {
        return nil, err
      }

      err = r.outport.Reformat(ctx, filename, nil)
      if err != nil {
        return nil, err
      }
    }
  }

  // router.go
  {
    filename := entity.GetControllerRouterFileName(*objCtrl)
    if !r.outport.IsFileExist(ctx, filename) {
      templateFile := r.outport.GetRouterTemplate(ctx, framework)

      err := r.outport.WriteFile(ctx, templateFile, filename, objDataCtrl)
      if err != nil {
        return nil, err
      }

      err = r.outport.Reformat(ctx, filename, nil)
      if err != nil {
        return nil, err
      }
    }
  }

  // inject inport to struct
  {
    templateCode := r.outport.GetRouterInportTemplate(ctx)

    templateWithData, err := r.outport.PrintTemplate(ctx, templateCode, objDataCtrl)
    if err != nil {
      return nil, err
    }

    bytes, err := objCtrl.InjectInportToStruct(templateWithData)
    if err != nil {
      return nil, err
    }

    // reformat outport.go
    err = r.outport.Reformat(ctx, entity.GetControllerRouterFileName(*objCtrl), bytes)
    if err != nil {
      return nil, err
    }
  }

  // inject router for register
  {
    templateCode := r.outport.GetRouterRegisterTemplate(ctx)

    templateWithData, err := r.outport.PrintTemplate(ctx, templateCode, objDataCtrl)
    if err != nil {
      return nil, err
    }

    bytes, err := objCtrl.InjectRouterBind(templateWithData)
    if err != nil {
      return nil, err
    }

    // reformat outport.go
    err = r.outport.Reformat(ctx, entity.GetControllerRouterFileName(*objCtrl), bytes)
    if err != nil {
      return nil, err
    }

  }

  return res, nil
}
