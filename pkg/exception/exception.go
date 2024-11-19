package exception

import (
	"fmt"
)

// Exceptions.
var (
	ErrUnauthorized        error = fmt.Errorf("%s", "Unauthorized")
	ErrNotFound            error = fmt.Errorf("%s", "Not found")
	ErrInternalServer      error = fmt.Errorf("%s", "Internal server error")
	ErrConflict            error = fmt.Errorf("%s", "Conflict")
	ErrUnprocessableEntity error = fmt.Errorf("%s", "Unprocessable entity")
	ErrBadRequest          error = fmt.Errorf("%s", "Bad request")
	ErrGatewayTimeout      error = fmt.Errorf("%s", "Gateway timeout")
	ErrTimeout             error = fmt.Errorf("%s", "Request time out")
	ErrLocked              error = fmt.Errorf("%s", "Locked")
	ErrForbidden           error = fmt.Errorf("%s", "Forbidden")
)
