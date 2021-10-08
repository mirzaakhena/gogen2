package templates

import _ "embed" // used for embed file

var (

	// InportFile ...
	//go:embed default/usecase/usecase_inport._go
	InportFile string

	// OutportFile ...
	//go:embed default/usecase/usecase_outport._go
	OutportFile string

	// InteractorFile ...
	//go:embed default/usecase/usecase_interactor._go
	InteractorFile string

	// TestFile ...
	//go:embed default/usecase/usecase_test._go
	TestFile string

	// LogFile ...
	//go:embed default/infrastructure/log/infra_log._go
	LogFile string

	// LogDefaultFile ...
	//go:embed default/infrastructure/log/infra_log_default._go
	LogDefaultFile string

	// EntityFile ...
	//go:embed default/domain/entity/entity._go
	EntityFile string
)

var (

	// RepositoryFile ...
	//go:embed default/domain/repository/repository._go
	RepositoryFile string

	// RepositoryInterfaceFile ...
	//go:embed default/domain/repository/repository_interface._go
	RepositoryInterfaceFile string

	// RepositoryInterfaceFindFile ...
	//go:embed default/domain/repository/repository_interface_find._go
	RepositoryInterfaceFindFile string

	// RepositoryInterfaceFindOneFile ...
	//go:embed default/domain/repository/repository_interface_findone._go
	RepositoryInterfaceFindOneFile string

	// RepositoryInterfaceRemoveFile ...
	//go:embed default/domain/repository/repository_interface_remove._go
	RepositoryInterfaceRemoveFile string

	// RepositoryInterfaceSaveFile ...
	//go:embed default/domain/repository/repository_interface_save._go
	RepositoryInterfaceSaveFile string
)

var (

	// RepoInjectInteractorFile ...
	//go:embed default/domain/repository/repository_inject._go
	RepoInjectInteractorFile string

	// RepoInjectInteractorFindFile ...
	//go:embed default/domain/repository/repository_inject_find._go
	RepoInjectInteractorFindFile string

	// RepoInjectInteractorFindOneFile ...
	//go:embed default/domain/repository/repository_inject_findone._go
	RepoInjectInteractorFindOneFile string

	// RepoInjectInteractorSaveFile ...
	//go:embed default/domain/repository/repository_inject_save._go
	RepoInjectInteractorSaveFile string

	// RepoInjectInteractorRemoveFile ...
	//go:embed default/domain/repository/repository_inject_remove._go
	RepoInjectInteractorRemoveFile string
)

var (

	// ApplicationFile ...
	//go:embed default/application/application._go
	ApplicationFile string

	// ApplicationErrorEnumFile ...
	//go:embed default/application/apperror/error_enum._go
	ApplicationErrorEnumFile string

	// ApplicationErrorFuncFile ...
	//go:embed default/application/apperror/error_func._go
	ApplicationErrorFuncFile string

	// ApplicationErrorTemplateFile ...
	//go:embed default/application/apperror/error_template._go
	ApplicationErrorTemplateFile string

	// ApplicationConstantTemplateFile ...
	//go:embed default/application/constant/constant._go
	ApplicationConstantTemplateFile string
)

var (

	// GatewayGormFile ...
	//go:embed default/gateway/gorm/impl._go
	GatewayGormFile string

	// GatewayGormInjectMethodFile ...
	//go:embed default/gateway/gorm/inject._go
	GatewayGormInjectMethodFile string
)

var(

	//go:embed default/application/registry/app-registry-gin._go
	RegistryGingonicFile string
)

var(

	//go:embed default/infrastructure/server/server-gin._go
	ServerGinGonicFile string
)

var (

	//go:embed default/controller/gingonic/handler-func._go
	ControllerGinGonicHandlerFuncFile string

	//go:embed default/controller/gingonic/interceptor._go
	ControllerGinGonicInterceptorFile string

	//go:embed default/controller/gingonic/response._go
	ControllerGinGonicResponseFile string

	//go:embed default/controller/gingonic/router-inport._go
	ControllerGinGonicRouterInportFile string

	//go:embed default/controller/gingonic/router-struct._go
	ControllerGinGonicRouterStructFile string

	//go:embed default/controller/gingonic/router-register._go
	ControllerGinGonicRouterRegisterFile string

)