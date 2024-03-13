package buffer

import "errors"

// ErrEndOfFile is the error returned when end of file or end of buffer is reached unexpectedly
// during reading or writing operations.
var ErrEndOfFile = errors.New("could not complete the operation as eof was reached unexpectedly")

// ErrInvalidByteOrder is the error returned when unknown endian id is provided in encoding/decoding of
// numeric datatypes
var ErrInvalidByteOrder = errors.New("could not parse the byteorder from the provided endianness id")

// ErrInvalidBool is the error returned when unknown boolbyte is provided in encoding/decoding of booleans
var ErrInvalidBool = errors.New("could not parse the boolbyte from the provided byte")

// ErrInvalidMagic is the error returned when the magic unconnected sequence could not be parsed
var ErrInvalidMagic = errors.New("could not parse the magic unconnected message sequence")
