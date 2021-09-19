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
