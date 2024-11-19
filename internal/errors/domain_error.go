package errors

import "errors"

var ErrUsernameUsed = errors.New("domain: username already used")

var ErrAuthentication = errors.New("domain: authentication error")