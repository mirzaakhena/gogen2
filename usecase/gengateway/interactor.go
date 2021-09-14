package gengateway

import "context"

//go:generate mockery --name Outport -output mocks/

type genGatewayInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase GenGateway
func NewUsecase(outputPort Outport) Inport {
	return &genGatewayInteractor{
		outport: outputPort,
	}
}

// Execute the usecase GenGateway
func (r *genGatewayInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	// code your usecase definition here ...

	return res, nil
}
