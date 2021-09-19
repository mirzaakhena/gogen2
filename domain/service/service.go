package service

import "context"

type CreateFolderIfNotExistService interface {
  CreateFolderIfNotExist(ctx context.Context, folderPath string) (bool, error)
}

type WriteFileIfNotExistService interface {
  WriteFileIfNotExist(ctx context.Context, templateFile, outputFilePath string, obj interface{}) (bool, error)
}

type WriteFileService interface {
  WriteFile(ctx context.Context, templateFile, outputFilePath string, obj interface{}) error
}

type ReformatService interface {
  Reformat(ctx context.Context, goFilename string, bytes []byte) error
}

type GetPackagePathService interface {
  GetPackagePath(ctx context.Context ) string
}

type IsFileExistService interface {
  IsFileExist(ctx context.Context, filepath string) bool
}

type PrintTemplateService interface {
  PrintTemplate(ctx context.Context, templateString string, x interface{}) (string, error)
}
