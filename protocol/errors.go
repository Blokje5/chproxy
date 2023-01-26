package protocol

import "errors"

var (
	ErrUnexpectedEOF = errors.New("unexpected EOF while reading from buffer")
)