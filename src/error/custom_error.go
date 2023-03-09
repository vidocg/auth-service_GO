package error

type AppError struct {
	Message       string
	HttpErrorCode int
	Error         error
}
