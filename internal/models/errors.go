package models

import (
	"errors"
)

// ErrNoRecord Custom error variable.
var ErrNoRecord = errors.New("models: no matching record found")
