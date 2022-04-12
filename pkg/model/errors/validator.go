package errors

var (
	//ErrInvokeValidatorEngineFailed ...
	ErrInvokeValidatorEngineFailed = NewStringError("Fail to invoke validator engine", 500)

	//ErrTypeAssertionFailed ...
	ErrTypeAssertionFailed = NewStringError("Fail to type assertion", 400)
)
