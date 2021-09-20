package prod

import (
	"context"
	"github.com/mirzaakhena/gogen2/infrastructure/templates"
)


// GetErrorTemplate ...
func (r prodGateway) GetErrorTemplate(ctx context.Context) (fun string, err string) {
	return templates.ApplicationErrorFuncFile, templates.ApplicationErrorEnumFile
}

// GetConstantTemplate ...
func (r prodGateway) GetConstantTemplate(ctx context.Context) string {
	return templates.ApplicationConstantTemplateFile
}

// GetApplicationTemplate ...
func (r prodGateway) GetApplicationTemplate(ctx context.Context) string {
	return templates.ApplicationFile
}
