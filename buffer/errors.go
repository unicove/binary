package buffer

import "errors"

// EOF_ERROR is the error returned when end of file or end of buffer is reached unexpectedly
// during reading or writing operations.
var EOF_ERROR = errors.New("could not complete the operation as eof was reached unexpectedly")

// CPE_ERROR is the error returned when unknown endian id is provided in encoding/decoding of
// numeric datatypes
var CPE_ERROR = errors.New("could not parse the byteorder from the provided endianness id")

// CPI_ERROR is the error returned when unknown ipaddress version is provided in encoding/decoding
// of IP addresses
var CPI_ERROR = errors.New("could not parse the ip address version from the provided version id")

// CPB_ERROR is the error returned when unknown boolbyte is provided in encoding/decoding of booleans
var CPB_ERROR = errors.New("could not parse the boolbyte from the provided byte")

// CPM_ERROR is the error returned when the magic unconnected sequence could not be parsed
var CPM_ERROR = errors.New("could not parse the magic unconnected message sequence")
