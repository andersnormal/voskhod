package server

import (
	"errors"
)

var (
	ErrReadyTimeout = errors.New("server: too long waiting for ready")
)
