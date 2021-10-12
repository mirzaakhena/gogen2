package ___genservice

import (
	"context"
)

// Inport of GenService
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase GenService
type InportRequest struct {
}

// InportResponse is response payload after running the usecase GenService
type InportResponse struct {
}
