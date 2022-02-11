package getting

import "errors"

// ErrNotFound is used when a key could not be found
var ErrNotFound = errors.New("getting: key not found")
