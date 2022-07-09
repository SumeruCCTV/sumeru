package errors

var (
	InvalidBody        = _error("invalid body")
	InvalidToken       = _error("invalid token")
	InvalidUsername    = _error("invalid username")
	InvalidCredentials = _error("invalid credentials")
	InvalidCameraName  = _error("invalid camera name")
	InvalidCameraAddr  = _error("invalid camera ip address")
	InvalidCameraPort  = _error("invalid camera port")
	InvalidCameraType  = _error("invalid camera type")

	ErrorLoggingIn          = _error("error logging in")
	ErrorLoggingOut         = _error("error logging out")
	ErrorRegisteringAccount = _error("error registering account")
	ErrorAddingCamera       = _error("error adding camera")

	NotFound        = _error("route not found")
	TooManyRequests = _error("too many requests")
)
