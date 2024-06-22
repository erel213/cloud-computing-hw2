package appError

type NotFoundError struct {
	Err error
}

func (e NotFoundError) Error() string {
	return e.Err.Error()
}

func (e NotFoundError) Code() int {
	return 404
}
