package herror

import "errors"

var (
	ErrNoChange         = errors.New("no change")
	ErrResourceNotFound = errors.New("resource not found")
	ErrRequired         = errors.New("required")
	ErrSessionExpired   = errors.New("session expired")
	ErrResourceDeleted  = errors.New("resource deleted (soft delete)")
)
