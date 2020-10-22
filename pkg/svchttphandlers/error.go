package svchttphandlers

const (
	// Please update the error code value code with the vcode of your service.

	// ErrNotFound is the error code returns for not found response.
	ErrNotFound = 1
	// ErrInternalCode is the error code returns for internal server error.
	ErrInternalCode = 2
	// ErrBadRequestCode is the error code returns for a bad request response.
	ErrBadRequestCode = 3
)

// ErrorResponse is the default error response with an error code and error_msg.
type ErrorResponse struct {
	VCode   int         `json:"vcode,omitempty"`
	Error   string      `json:"error,omitempty"`
	Details interface{} `json:"details,omitempty"`
}

// GetErrorResponse converts any application error to response error.
func getErrorResponse(err error) ErrorResponse {
	errCode := ErrInternalCode
	switch err {
	// case ErrSth1:
	// 	errCode = ErrBadRequestCode

	// case ErrSth2:
	// 	errCode = ErrNotFound

	default:
		errCode = ErrInternalCode
	}

	return errResponse(errCode, err)
}

var errorMessage = map[int]string{
	ErrNotFound:       "not found",
	ErrBadRequestCode: "bad request",
	ErrInternalCode:   "internal server error",
}

func errResponse(errCode int, err error) ErrorResponse {
	return ErrorResponse{
		VCode:   errCode,
		Error:   errorMessage[errCode],
		Details: err.Error(),
	}
}
