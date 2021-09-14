package service

import "context"

type LogActionInterface interface {



  GetLogRootFolderName(ctx context.Context) string
  GetLogInterfaceFileName(ctx context.Context) string
  GetLogImplementationFileName(ctx context.Context) string
  GetLogInterfaceTemplate(ctx context.Context) string
  GetLogImplementationTemplate(ctx context.Context) string

  CreateFolderIfNotExistService
  WriteFileIfNotExistService

}
