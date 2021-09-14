package genregistry

import "context"

//go:generate mockery --name Outport -output mocks/

type genRegistryInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase GenRegistry
func NewUsecase(outputPort Outport) Inport {
	return &genRegistryInteractor{
		outport: outputPort,
	}
}

// Execute the usecase GenRegistry
func (r *genRegistryInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	// code your usecase definition here ...

	return res, nil
}
