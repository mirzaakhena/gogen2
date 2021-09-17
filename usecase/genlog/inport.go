package genlog

import (
	"context"
)

// Inport of GenLog
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase GenLog
type InportRequest struct {
}

// InportResponse is response payload after running the usecase GenLog
type InportResponse struct {
}
