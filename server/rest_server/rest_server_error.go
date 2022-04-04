package rest_server

type ServerError struct {
	error string
	code  int
}

func (se ServerError) Error() string {
	return se.error
}

func (se ServerError) Code() int {
	return se.code
}
