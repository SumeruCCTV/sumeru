package errors

type _errorType string

func (e _errorType) Error() string {
	return string(e)
}

func _error(err string) error {
	return New(_errorType(err))
}
