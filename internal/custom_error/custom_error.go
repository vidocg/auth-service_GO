package custom_error

type AppError struct {
	Message       string
	HttpErrorCode int
	Error         error
}
