package errors

import (
	"errors"

	"connectrpc.com/connect"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func New(err error) error {
	if err == nil {
		return status.Error(codes.Internal, "unknown error")
	}
	if e, ok := status.FromError(err); ok {
		return e.Err()
	}
	return NewInternalError(err.Error())
}

func NewInternalError(message string) error {
	return connect.NewError(connect.CodeInternal, errors.New(message))
}

func NewNotFound(message string) error {
	return connect.NewError(connect.CodeNotFound, errors.New(message))
}

func NewBadRequest(message string) error {
	return connect.NewError(connect.CodeInvalidArgument, errors.New(message))
}

func NewAlreadyExists(message string) error {
	return connect.NewError(connect.CodeAlreadyExists, errors.New(message))
}

func NewForbidden(message string) error {
	return connect.NewError(connect.CodePermissionDenied, errors.New(message))
}

func NewUnauthenticated(message string) error {
	return connect.NewError(connect.CodeUnauthenticated, errors.New(message))
}

func NewUnauthorized(message string) error {
	return NewUnauthenticated(message)
}

func NewNotImplemented(message string) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New(message))
}
