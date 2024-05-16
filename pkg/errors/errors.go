package errors

import "errors"

var ErrRecordNotFound = errors.New("models: no matching record found")
var ErrNoRowUpdated = errors.New("models: no row updated")
var ErrNoRowDeleted = errors.New("models: no row deleted")
