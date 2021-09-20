package prod

import (
	"context"
	"github.com/mirzaakhena/gogen2/domain/vo"
	"github.com/mirzaakhena/gogen2/infrastructure/templates"
)

func (r *prodGateway) GetRepositoryTemplate(ctx context.Context) string {
	return templates.RepositoryFile
}

func (r *prodGateway) GetRepositoryFunctionTemplate(ctx context.Context, repoName vo.Naming) (string, error) {

	if repoName.HasOneOfThisPrefix("save", "create", "add", "update") {
		return templates.RepositoryInterfaceSaveFile, nil
	}

	if repoName.HasOneOfThisPrefix("findone", "findfirst", "findlast", "getone") {
		return templates.RepositoryInterfaceFindOneFile, nil
	}

	if repoName.HasOneOfThisPrefix("find", "get") {
		return templates.RepositoryInterfaceFindFile, nil
	}

	if repoName.HasOneOfThisPrefix("remove", "delete") {
		return templates.RepositoryInterfaceRemoveFile, nil
	}

	return templates.RepositoryInterfaceFile, nil
}

func (r *prodGateway) GetInteractorRepoCallTemplate(ctx context.Context, repoName vo.Naming) (string, error) {

	if repoName.HasOneOfThisPrefix("findone", "findfirst", "findlast", "getone") { //
		return templates.RepoInjectInteractorFindOneFile, nil

	}

	if repoName.HasOneOfThisPrefix("find", "get") {
		return templates.RepoInjectInteractorFindFile, nil

	}

	if repoName.HasOneOfThisPrefix("remove", "delete") {
		return templates.RepoInjectInteractorRemoveFile, nil

	}

	if repoName.HasOneOfThisPrefix("save", "create", "add", "update") {
		return templates.RepoInjectInteractorSaveFile, nil

	}

	return templates.RepoInjectInteractorFile, nil

}
