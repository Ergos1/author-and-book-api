package author

import "errors"

var ErrAuthorDuplicate = errors.New("author is not unique")
var ErrAuthorNotFound = errors.New("author not found")
