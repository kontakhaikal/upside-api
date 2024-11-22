package errors

import "errors"

var ErrUsernameUsed = errors.New("domain: username already used")

var ErrAuthentication = errors.New("domain: authentication error")

var ErrAlreadyJoinedSide = errors.New("domain already joined side")

var ErrSideNotFound = errors.New("domain: side not found")

var ErrNotAMember = errors.New("domain: author not a member of the side")