package http_errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	InvalidNameOrPassword = "invalid email or password"
	InvalidName           = "the name is already taken"
	InsufficientFunds     = "insufficient funds"
)

func HandleError(err error) (string, int) {
	st, ok := status.FromError(err)
	if !ok {
		return "", 500
	}
	mes := st.Message()
	var code int
	switch st.Code() {
	case codes.InvalidArgument:
		code = 400
	case codes.Unauthenticated:
		code = 401
	case codes.PermissionDenied:
		code = 403
	case codes.NotFound:
		code = 404
	default:
		code = 500
	}
	return mes, code
}
