package appError

import (
	"fmt"
	"net/http"
)

type HttpError struct {
	StatusCode int
	Err        error
}

func (r *HttpError) Error() string {
	return r.Err.Error()
}

var ErrInternalServer = &HttpError{
	Err:        fmt.Errorf("unhandled error"),
	StatusCode: http.StatusInternalServerError,
}
