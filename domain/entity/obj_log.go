package entity

import (
  "context"
  "github.com/mirzaakhena/gogen2/domain/service"
)

type ObjLog struct {
}

func (o ObjLog) Construct(ctx context.Context, logConstructAction service.LogActionInterface) error {

  {
    err := logConstructAction.CreateFolderIfNotExist(ctx, logConstructAction.GetLogRootFolderName(ctx))
    if err != nil {
      return err
    }
  }

  {
    logTemplateFile := logConstructAction.GetLogInterfaceTemplate(ctx)
    outputFile := logConstructAction.GetLogInterfaceFileName(ctx)
    err := logConstructAction.WriteFileIfNotExist(ctx, logTemplateFile, outputFile, struct{}{})
    if err != nil {
      return err
    }
  }

  {
    logImplTemplateFile := logConstructAction.GetLogImplementationTemplate(ctx)
    outputFile := logConstructAction.GetLogImplementationFileName(ctx)
    err := logConstructAction.WriteFileIfNotExist(ctx, logImplTemplateFile, outputFile, struct{}{})
    if err != nil {
      return err
    }
  }

  return nil
}

//func (o ObjLog) GetLogRootFolderName(ctx context.Context) string {
//  return fmt.Sprintf("infrastructure/log")
//}
//
//func (o ObjLog) GetLogInterfaceFileName(ctx context.Context) string {
//  return fmt.Sprintf("infrastructure/log/log.go")
//}
//
//func (o ObjLog) GetLogImplementationFileName(ctx context.Context) string {
//  return fmt.Sprintf("infrastructure/log/log_default.go")
//}

//{
//err := r.outport.CreateFolderIfNotExist(ctx, "infrastructure/log")
//if err != nil {
//return nil, err
//}
//
//{
//logTemplateFile := templates.LogFile
//outputFile := fmt.Sprintf("infrastructure/log/log.go")
//err = r.outport.WriteFileIfNotExist(ctx, logTemplateFile, outputFile, struct{}{})
//if err != nil {
//return nil, err
//}
//}
//
//{
//logImplTemplateFile := templates.LogDefaultFile
//outputFile := fmt.Sprintf("infrastructure/log/log_default.go")
//err = r.outport.WriteFileIfNotExist(ctx, logImplTemplateFile, outputFile, struct{}{})
//if err != nil {
//return nil, err
//}
//}
//}
