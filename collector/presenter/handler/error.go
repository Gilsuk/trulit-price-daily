package handler

import "errors"

var ErrInvalidDateFormat error = errors.New("invalid date format")
var ErrInvalidJsonSyntax error = errors.New("invalid json syntax")
