package rest_client

type ClientError struct {
	error string
	code  int
}

func (ce ClientError) Error() string {
	return ce.error
}

func (ce ClientError) Code() int {
	return ce.code
}
