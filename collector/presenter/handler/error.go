package handler

import "errors"

var ErrInvalidDateFormat error = errors.New("invalid date format. the format should be yyyy-MM-dd")
var ErrInvalidJsonSyntax error = errors.New("invalid json syntax")
