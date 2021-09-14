package genvaluestring

import "context"

//go:generate mockery --name Outport -output mocks/

type genValueStringInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase GenValueString
func NewUsecase(outputPort Outport) Inport {
	return &genValueStringInteractor{
		outport: outputPort,
	}
}

// Execute the usecase GenValueString
func (r *genValueStringInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	// code your usecase definition here ...

	return res, nil
}
