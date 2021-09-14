package service

import "context"

type CreateFolderIfNotExistService interface {
  CreateFolderIfNotExist(ctx context.Context, folderPath string) error
}

type WriteFileIfNotExistService interface {
  WriteFileIfNotExist(ctx context.Context, templateFile, outputFilePath string, obj interface{}) error
}

type WriteFileService interface {
  WriteFile(ctx context.Context, templateFile, outputFilePath string, obj interface{}) error
}

type ReformatService interface {
  Reformat(ctx context.Context, goFilename string) error
}

type GetPackagePathService interface {
  GetPackagePath(ctx context.Context ) string
}

type IsFileExistService interface {
  IsFileExist(ctx context.Context, filepath string) bool
}
