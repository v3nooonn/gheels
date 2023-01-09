package errorx

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"net/http"
)

type Err struct {
	CodeF    int    `json:"code"`
	StatusF  string `json:"status"`
	MessageF string `json:"message"`
}

func (r *Err) Code() int {
	return r.CodeF
}

func (r *Err) Status() string {
	return r.StatusF
}

func (r *Err) Message() string {
	return r.MessageF
}

func (r *Err) Error() string {
	return fmt.Sprintf("code: %v, status: %s, message: %s", r.Code(), r.Status(), r.Message())
}

func newErr(c int, s string, m string) *Err {
	return &Err{
		CodeF:    c,
		StatusF:  s,
		MessageF: m,
	}
}

func BadRequest(m string) error {
	return newErr(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), m)
}

func Unauthorized(m string) *Err {
	return newErr(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), m)
}

func Forbidden(m string) *Err {
	return newErr(http.StatusForbidden, http.StatusText(http.StatusForbidden), m)
}

func NotFound(m string) *Err {
	return newErr(http.StatusNotFound, http.StatusText(http.StatusNotFound), m)
}

func Internal(m string) *Err {
	return newErr(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), m)
}

func AlreadyExists(m string) error {
	return newErr(http.StatusConflict, http.StatusText(http.StatusConflict), m)
}

// HTTPStatusFromCode converts a gRPC error code into the corresponding HTTP response status.
// See: https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
func HTTPStatusFromCode(code codes.Code) uint32 {
	switch code {
	case codes.OK:
		return http.StatusOK
	case codes.InvalidArgument, codes.FailedPrecondition, codes.OutOfRange:
		return http.StatusBadRequest
	case codes.Unauthenticated:
		return http.StatusUnauthorized
	case codes.PermissionDenied:
		return http.StatusForbidden
	case codes.NotFound:
		return http.StatusNotFound
	case codes.AlreadyExists, codes.Aborted:
		return http.StatusConflict
	case codes.ResourceExhausted:
		return http.StatusTooManyRequests
	case codes.Canceled:
		return 499
	case codes.Unknown, codes.Internal, codes.DataLoss:
		return http.StatusInternalServerError
	case codes.Unimplemented:
		return http.StatusNotImplemented
	case codes.Unavailable:
		return http.StatusServiceUnavailable
	case codes.DeadlineExceeded:
		return http.StatusGatewayTimeout
	}

	return http.StatusInternalServerError
}
