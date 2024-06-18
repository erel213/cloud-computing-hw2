package appError

type AppError interface {
	Error() string
	Code() int
}
