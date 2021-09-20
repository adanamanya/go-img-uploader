package errors

import "errors"

// New returns an error that formats as the given text.
var New = errors.New

// List of known errors
var (
	ErrInvalidFileType = errors.New("file type not supported")
	ErrFailedInit      = errors.New("initialization failed, check credentials")
	ErrStorageinit     = errors.New("failed init storage")
	ErrParsingFile     = errors.New("error parsing file")
	ErrNotFound        = errors.New("url not found")
	ErrNotAllowed      = errors.New("method not allowed")
	ErrTooLarge        = errors.New("content must be smaller than 8mb")
	ErrFalseContent    = errors.New("content must be image")
	ErrFailedUpload    = errors.New("failed upload image, check your connection")
	ErrUnauthorized    = errors.New("auth token did not match")
)
