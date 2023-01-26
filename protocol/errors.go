package protocol

import "errors"

var (
	ErrUnexpectedEOF = errors.New("unexpected EOF while reading from buffer")
	ErrStringExceedsMaximum = errors.New("string size in TCP message exceeds 10mb maximum")
)