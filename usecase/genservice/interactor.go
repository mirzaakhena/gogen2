package genservice

import "context"

//go:generate mockery --name Outport -output mocks/

type genServiceInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase GenService
func NewUsecase(outputPort Outport) Inport {
	return &genServiceInteractor{
		outport: outputPort,
	}
}

// Execute the usecase GenService
func (r *genServiceInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	// code your usecase definition here ...

	return res, nil
}
