package asn

import "errors"

// ErrInvalidInput is returned when parsing string input fails.
var ErrInvalidInput = errors.New("invalid input")

// ErrOutOf2ByteRange is returned when parsing string input fails because a required 16 bit value
// exceeds the allowable range.
var ErrOutOf2ByteRange = errors.New("value out of range (0-65535)")

// ErrOutOf2ByteRange is returned when parsing string input fails because a required 32 bit value
// exceeds the allowable range.
var ErrOutOf4ByteRange = errors.New("value out of range (0-4294967295)")
