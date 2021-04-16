package utils

type BusinessError struct {
	Code int
	Err  string
}

func NewBusinessError(code int, err string) error {
	return &BusinessError{code, err}
}

func (e *BusinessError) Error() string {
	var errorMeg string
	if e != nil {
		errorMeg = e.Err
	}
	return errorMeg
}
