package appError

type AuthenticationError struct {
	Err error
}

func (e AuthenticationError) Error() string {
	return e.Err.Error()
}

func (e AuthenticationError) Code() int {
	return 401
}
