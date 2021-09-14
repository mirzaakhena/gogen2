package gencontroller

import "context"

//go:generate mockery --name Outport -output mocks/

type genControllerInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase GenController
func NewUsecase(outputPort Outport) Inport {
	return &genControllerInteractor{
		outport: outputPort,
	}
}

// Execute the usecase GenController
func (r *genControllerInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	// code your usecase definition here ...

	return res, nil
}
