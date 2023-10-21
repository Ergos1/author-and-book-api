package handlers

import "errors"

var ErrInvalidId = errors.New("invalid id")
var ErrAuthorNotFound = errors.New("author not found")
var ErrBookNotFound = errors.New("book not found")
var ErrInvalidRequestBody = errors.New("invalid request body")
