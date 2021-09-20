package prod

import (
	"context"
	"github.com/mirzaakhena/gogen2/infrastructure/templates"
)

func (r prodGateway) GetErrorTemplate(ctx context.Context) (fun string, err string) {
	return templates.ApplicationErrorFuncFile, templates.ApplicationErrorEnumFile
}

func (r prodGateway) GetConstantTemplate(ctx context.Context) string {
	return templates.ApplicationConstantTemplateFile
}

func (r prodGateway) GetApplicationTemplate(ctx context.Context) string {
	return templates.ApplicationFile
}
