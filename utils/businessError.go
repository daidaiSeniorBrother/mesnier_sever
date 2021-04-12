package utils

type BusinessError struct {
	Code int
	Err  string
}

func NewBusinessError(code int, err string) error {
	return &BusinessError{code, err}
}

func (e *BusinessError) Error() string {
	return e.Err
}
