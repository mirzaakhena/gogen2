package prod

import (
	"context"
	"github.com/mirzaakhena/gogen2/domain/vo"
	"github.com/mirzaakhena/gogen2/infrastructure/templates"
)

// GetRepositoryTemplate ...
func (r *prodGateway) GetRepositoryTemplate(ctx context.Context) string {
	return templates.ReadFile("domain/repository/repository._go")
}

// GetRepositoryFunctionTemplate ...
func (r *prodGateway) GetRepositoryFunctionTemplate(ctx context.Context, repoName vo.Naming) (string, error) {

	if repoName.HasOneOfThisPrefix("save", "create", "add", "update") {
		return templates.ReadFile("domain/repository/~repository_interface_save._go"), nil
	}

	if repoName.HasOneOfThisPrefix("findone", "findfirst", "findlast", "getone") {
		return templates.ReadFile("domain/repository/~repository_interface_findone._go"), nil
	}

	if repoName.HasOneOfThisPrefix("find", "get") {
		return templates.ReadFile("domain/repository/~repository_interface_find._go"), nil
	}

	if repoName.HasOneOfThisPrefix("remove", "delete") {
		return templates.ReadFile("domain/repository/~repository_interface_remove._go"), nil
	}

	return templates.ReadFile("domain/repository/~repository_interface._go"), nil
}

// GetInteractorRepoCallTemplate ...
func (r *prodGateway) GetInteractorRepoCallTemplate(ctx context.Context, repoName vo.Naming) (string, error) {

	if repoName.HasOneOfThisPrefix("findone", "findfirst", "findlast", "getone") { //
		return templates.ReadFile("domain/repository/~interactor_inject_findone._go"), nil
	}

	if repoName.HasOneOfThisPrefix("find", "get") {
		return templates.ReadFile("domain/repository/~interactor_inject_find._go"), nil
	}

	if repoName.HasOneOfThisPrefix("remove", "delete") {
		return templates.ReadFile("domain/repository/~interactor_inject_remove._go"), nil
	}

	if repoName.HasOneOfThisPrefix("save", "create", "add", "update") {
		return templates.ReadFile("domain/repository/~interactor_inject_save._go"), nil

	}

	return templates.ReadFile("domain/repository/~interactor_inject._go"), nil

}
