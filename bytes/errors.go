package bytes

// EOF_ERROR is the error returned when end of file or end of buffer is reached unexpectedly
// during reading or writing operations.
const EOF_ERROR = "could not complete the operation as eof was reached unexpectedly"

// CPB_ERROR is the error returned when unknown endian id is provided in encoding/decoding of
// numeric datatypes
const CPB_ERROR = "could not parse the byteorder from the provided endianness id"
