
// 错误码
package errno

import "google.golang.org/grpc/codes"

type Error struct {
	Code    codes.Code
	Message string
}

var (
	ErrInvalidArgument = &Error{Code: codes.InvalidArgument, Message: "Invalid argument"}
	ErrInternal        = &Error{Code: codes.Internal, Message: "Internal error"}
	ErrPaymentFailed   = &Error{Code: codes.Internal, Message: "Payment processing failed"}
)
