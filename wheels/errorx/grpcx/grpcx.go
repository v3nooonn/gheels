package grpcx

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// BadRequest indicates client specified an invalid argument.
func BadRequest(format string, a ...interface{}) error {
	return status.Errorf(codes.InvalidArgument, format, a...)
}

// NotFound means some requested entity (e.g., file or directory) was
// not found.
func NotFound(format string, a ...interface{}) error {
	return status.Errorf(codes.NotFound, format, a...)
}

// AlreadyExists means an attempt to create an entity failed because one
// already exists.
func AlreadyExists(format string, a ...interface{}) error {
	return status.Errorf(codes.AlreadyExists, format, a...)
}

// PermissionDenied indicates the caller does not have permission to
// execute the specified operation.
func PermissionDenied(format string, a ...interface{}) error {
	return status.Errorf(codes.PermissionDenied, format, a...)
}

// Aborted indicates the operation was aborted, typically due to a
// concurrency issue like sequencer check failures, transaction aborts, etc.
func Aborted(format string, a ...interface{}) error {
	return status.Errorf(codes.Aborted, format, a...)
}

// OutOfRange means operation was attempted past the valid range.
// E.g., seeking or reading past end of file.
func OutOfRange(format string, a ...interface{}) error {
	return status.Errorf(codes.OutOfRange, format, a...)
}

// Internal errors. Means some invariants expected by underlying
// system has been broken. If you see one of these errors,
// something is very broken.
func Internal(format string, a ...interface{}) error {
	return status.Errorf(codes.Internal, format, a...)
}

// DeadlineExceeded means operation expired before completion.
// For operations that change the state of the system, this error may be
// returned even if the operation has completed successfully. For
// example, a successful response from a server could have been delayed
// long enough for the deadline to expire.
func DeadlineExceeded(format string, a ...interface{}) error {
	return status.Errorf(codes.DeadlineExceeded, format, a...)
}

// Unimplemented indicates operation is not implemented or not
// supported/enabled in this service.
func Unimplemented(format string, a ...interface{}) error {
	return status.Errorf(codes.Unimplemented, format, a...)
}
