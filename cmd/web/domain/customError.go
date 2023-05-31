package domain

type CustomErrorWrapper struct {
	Message string `json:"message"`
	Code    int    `json:"-"`
	Err     error  `json:"-"`
}
func NewErrorWrapper(code int, err error, message string) error {
	return CustomErrorWrapper{
		Message: message,
		Code: code,
		Err: err,
	}
}
func (err CustomErrorWrapper) Error() string {
	// guard against panics
	if err.Err != nil {
		return err.Err.Error()
	}
	return err.Message
}
func (err CustomErrorWrapper) Unwrap() error {
	return err.Err // Returns inner error
}
