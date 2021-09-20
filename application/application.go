package application

// RegistryContract ...
type RegistryContract interface {
	SetupController()
	RunApplication()
}

// Run ...
func Run(rv RegistryContract) {
	if rv != nil {
		rv.SetupController()
		rv.RunApplication()
	}
}
