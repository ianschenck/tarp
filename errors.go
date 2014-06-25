package tarp

import (
	"errors"
)

var (
	ErrSemverViolation = errors.New("name violates semver")
	ErrBadName         = errors.New("bad name")
)
