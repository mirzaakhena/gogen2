package genentity

import "context"

//go:generate mockery --name Outport -output mocks/

type genEntityInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase GenEntity
func NewUsecase(outputPort Outport) Inport {
	return &genEntityInteractor{
		outport: outputPort,
	}
}

// Execute the usecase GenEntity
func (r *genEntityInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	// code your usecase definition here ...

	return res, nil
}
