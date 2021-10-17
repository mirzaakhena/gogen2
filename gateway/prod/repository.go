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

// GetRepoInterfaceTemplate ...
func (r *prodGateway) GetRepoInterfaceTemplate(ctx context.Context, repoName vo.Naming) string {

  if repoName.HasOneOfThisPrefix("save", "create", "add", "update") {
    return templates.ReadFile("domain/repository/~repository_interface_save._go")
  }

  if repoName.HasOneOfThisPrefix("findone", "findfirst", "findlast", "getone") {
    return templates.ReadFile("domain/repository/~repository_interface_findone._go")
  }

  if repoName.HasOneOfThisPrefix("find", "get") {
    return templates.ReadFile("domain/repository/~repository_interface_find._go")
  }

  if repoName.HasOneOfThisPrefix("remove", "delete") {
    return templates.ReadFile("domain/repository/~repository_interface_remove._go")
  }

  return templates.ReadFile("domain/repository/~repository_interface._go")
}

// GetRepoInjectTemplate ...
func (r *prodGateway) GetRepoInjectTemplate(ctx context.Context, repoName vo.Naming) string {

  if repoName.HasOneOfThisPrefix("findone", "findfirst", "findlast", "getone") { //
    return templates.ReadFile("domain/repository/~interactor_inject_findone._go")
  }

  if repoName.HasOneOfThisPrefix("find", "get") {
    return templates.ReadFile("domain/repository/~interactor_inject_find._go")
  }

  if repoName.HasOneOfThisPrefix("remove", "delete") {
    return templates.ReadFile("domain/repository/~interactor_inject_remove._go")
  }

  if repoName.HasOneOfThisPrefix("save", "create", "add", "update") {
    return templates.ReadFile("domain/repository/~interactor_inject_save._go")

  }

  return templates.ReadFile("domain/repository/~interactor_inject._go")

}
