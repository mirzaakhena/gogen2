package service

import (
	"context"
	"github.com/mirzaakhena/gogen2/domain/entity"
)

// LogActionInterface ...
type LogActionInterface interface {
	CreateFolderIfNotExistService
	WriteFileIfNotExistService
	GetLogInterfaceTemplate(ctx context.Context) string
	GetLogImplementationTemplate(ctx context.Context) string
}

// ConstructLog ...
func ConstructLog(ctx context.Context, action LogActionInterface) error {

	{
		_, err := action.CreateFolderIfNotExist(ctx, entity.GetLogRootFolderName())
		if err != nil {
			return err
		}
	}

	{
		logTemplateFile := action.GetLogInterfaceTemplate(ctx)
		outputFile := entity.GetLogInterfaceFileName()
		_, err := action.WriteFileIfNotExist(ctx, logTemplateFile, outputFile, struct{}{})
		if err != nil {
			return err
		}
	}

	{
		logImplTemplateFile := action.GetLogImplementationTemplate(ctx)
		outputFile := entity.GetLogImplementationFileName()
		_, err := action.WriteFileIfNotExist(ctx, logImplTemplateFile, outputFile, struct{}{})
		if err != nil {
			return err
		}
	}

	return nil
}
