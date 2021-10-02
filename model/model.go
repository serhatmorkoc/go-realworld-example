package model

import "github.com/pkg/errors"

const ISO8601 = "2006-01-02T15:04:05-0700Z"

var (
	ErrOperationFailed = errors.New("operation failed")
)