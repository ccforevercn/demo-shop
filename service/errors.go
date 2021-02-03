package service

type ErrorService struct {
	error
	code int
	message  string
}

func (errorService *ErrorService) Error() (string, int) {
	return errorService.message, errorService.code
}

func (errorService *ErrorService) SetError(code int, message string) *ErrorService {
	return &ErrorService{
		code: code,
		message: message,
	}
}
