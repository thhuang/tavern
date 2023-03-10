package error

type ErrorCode int

const (
	ErrorCodeUnknown        ErrorCode = 13703
	ErrorCodeNotImplemented ErrorCode = 13704
	ErrorCodeBadRequest     ErrorCode = 13705
)

type ErrorResponse struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
}
