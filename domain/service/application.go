package service

import (
	"context"
	"fmt"
)

// ApplicationActionInterface ...
type ApplicationActionInterface interface {
	CreateFolderIfNotExistService
	WriteFileIfNotExistService
	GetErrorTemplate(ctx context.Context) (string, string)
	GetConstantTemplate(ctx context.Context) string
	GetApplicationTemplate(ctx context.Context) string
}

// ConstructApplication ...
func ConstructApplication(ctx context.Context, action ApplicationActionInterface) error {

	{
		_, _ = action.CreateFolderIfNotExist(ctx, "application/apperror")
		_, _ = action.CreateFolderIfNotExist(ctx, "application/constant")
		_, _ = action.CreateFolderIfNotExist(ctx, "application/registry")
	}

	errorEnum, errorFunc := action.GetErrorTemplate(ctx)
	{
		outputFile := fmt.Sprintf("application/apperror/error_enum.go")
		_, err := action.WriteFileIfNotExist(ctx, errorEnum, outputFile, struct{}{})
		if err != nil {
			return err
		}
	}

	{
		outputFile := fmt.Sprintf("application/apperror/error_func.go")
		_, err := action.WriteFileIfNotExist(ctx, errorFunc, outputFile, struct{}{})
		if err != nil {
			return err
		}
	}

	{
		appTemplateFile := action.GetConstantTemplate(ctx)
		outputFile := fmt.Sprintf("application/constant/constant.go")
		_, err := action.WriteFileIfNotExist(ctx, appTemplateFile, outputFile, struct{}{})
		if err != nil {
			return err
		}
	}

	{
		appTemplateFile := action.GetApplicationTemplate(ctx)
		outputFile := fmt.Sprintf("application/application.go")
		_, err := action.WriteFileIfNotExist(ctx, appTemplateFile, outputFile, struct{}{})
		if err != nil {
			return err
		}
	}

	return nil
}
