package templates

import _ "embed"

//go:embed default/usecase/usecase_inport._go
var InportFile string

//go:embed default/usecase/usecase_outport._go
var OutportFile string

//go:embed default/usecase/usecase_interactor._go
var InteractorFile string

//go:embed default/usecase/usecase_test._go
var TestFile string

//go:embed default/infrastructure/log/infra_log._go
var LogFile string

//go:embed default/infrastructure/log/infra_log_default._go
var LogDefaultFile string

//go:embed default/domain/entity/entity._go
var EntityFile string

var (
	//go:embed default/domain/repository/repository._go
	RepositoryFile string

	//go:embed default/domain/repository/repository_interface._go
	RepositoryInterfaceFile string

	//go:embed default/domain/repository/repository_interface_find._go
	RepositoryInterfaceFindFile string

	//go:embed default/domain/repository/repository_interface_findone._go
	RepositoryInterfaceFindOneFile string

	//go:embed default/domain/repository/repository_interface_remove._go
	RepositoryInterfaceRemoveFile string

	//go:embed default/domain/repository/repository_interface_save._go
	RepositoryInterfaceSaveFile string
)

var (
	//go:embed default/domain/repository/repository_inject._go
	RepoInjectInteractorFile string

	//go:embed default/domain/repository/repository_inject_find._go
	RepoInjectInteractorFindFile string

	//go:embed default/domain/repository/repository_inject_findone._go
	RepoInjectInteractorFindOneFile string

	//go:embed default/domain/repository/repository_inject_save._go
	RepoInjectInteractorSaveFile string

	//go:embed default/domain/repository/repository_inject_remove._go
	RepoInjectInteractorRemoveFile string
)

var (
	//go:embed default/application/application._go
	ApplicationFile string

	//go:embed default/application/apperror/error_enum._go
	ApplicationErrorEnumFile string

	//go:embed default/application/apperror/error_func._go
	ApplicationErrorFuncFile string

	//go:embed default/application/apperror/error_template._go
	ApplicationErrorTemplateFile string

	//go:embed default/application/constant/constant._go
	ApplicationConstantTemplateFile string
)
