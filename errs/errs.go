package errs

type AppError struct {
	Code      int
	Messesage string
}

func (e AppError) Error() string {
	return e.Messesage
}
