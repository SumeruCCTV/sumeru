package errors

var (
	InvalidBody        = _error("invalid body")
	InvalidToken       = _error("invalid token")
	InvalidUsername    = _error("invalid username")
	InvalidCredentials = _error("invalid credentials")

	ErrorLoggingIn          = _error("error logging in")
	ErrorLoggingOut         = _error("error logging out")
	ErrorRegisteringAccount = _error("error registering account")

	NotFound        = _error("route not found")
	TooManyRequests = _error("too many requests")
)
