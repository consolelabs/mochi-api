package errors

var (
	// ErrRecordNotFound not found record by id
	ErrRecordNotFound = NewStringError("record not found", 404)

	// ErrInternalError internal error
	ErrInternalError = NewStringError("internal error", 500)

	// ErrBadRequest internal error
	ErrBadRequest = NewStringError("bad request", 400)
)
