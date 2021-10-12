package ___genvalueobject

import "context"

//go:generate mockery --name Outport -output mocks/

type genValueObjectInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase GenValueObject
func NewUsecase(outputPort Outport) Inport {
	return &genValueObjectInteractor{
		outport: outputPort,
	}
}

// Execute the usecase GenValueObject
func (r *genValueObjectInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	// code your usecase definition here ...

	return res, nil
}
