package db

import "github.com/pkg/errors"

var (
	errUnSupportedDriver = errors.New("errUnSupportedDriver")
	errPingDatabase = errors.New("errPingDatabase")
)
