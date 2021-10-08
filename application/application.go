package application

import "github.com/mirzaakhena/gogen2/controller"

// RegistryContract ...
type RegistryContract interface {
	controller.Controller
	RunApplication()
}

// Run ...
func Run(rv RegistryContract) {
	if rv != nil {
		rv.RegisterRouter()
		rv.RunApplication()
	}
}
