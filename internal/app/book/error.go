package book

import "errors"

var ErrBookDuplicate = errors.New("book is not unique")
var ErrBookNotFound = errors.New("book not found")
