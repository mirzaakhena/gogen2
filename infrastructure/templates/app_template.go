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

//go:embed default/domain/repository/repository._go
var RepositoryFile string

//go:embed default/infrastructure/log/infra_log._go
var LogFile string

//go:embed default/infrastructure/log/infra_log_default._go
var LogDefaultFile string

//go:embed default/domain/entity/entity._go
var EntityFile string


