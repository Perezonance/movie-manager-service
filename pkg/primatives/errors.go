package primatives

import (
	"github.com/pkg/errors"
)

var(
	ErrUserNotFound = errors.New("User not found in table")
	ErrPostNotFound = error.New("Post not found in table")
)
