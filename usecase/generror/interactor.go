package generror

import "context"

//go:generate mockery --name Outport -output mocks/

type genErrorInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase GenError
func NewUsecase(outputPort Outport) Inport {
	return &genErrorInteractor{
		outport: outputPort,
	}
}

// Execute the usecase GenError
func (r *genErrorInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	// code your usecase definition here ...

	return res, nil
}
