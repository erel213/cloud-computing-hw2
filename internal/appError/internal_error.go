package appError

type InternalError struct {
	Err error
}

func (e InternalError) Error() string {
	return e.Err.Error()
}

func (e InternalError) Code() int {
	return 500
}
