package bytes

// EOF_ERROR is the error returned when end of file or end of buffer is reached unexpectedly
// during reading or writing operations.
const EOF_ERROR = "could not complete the operation as eof was reached unexpectedly"

// Buffer represents a fixed size buffer for reading and writing various Minecraft Datatypes over the wire.
// It's capacity is fixed and cannot be dynamically increased, thus increasing performance. It is useful for
// those scenarios where the exact size or the max size of the data you want to receive is known.
type Buffer struct {
	slice       []byte
	cap         int
	readOffset  int
	writeOffset int
}

// Creates and returns a new Buffer from the provided slice.
func New(slice []byte) *Buffer {
	return &Buffer{
		slice:       slice,
		cap:         cap(slice) - 1,
		readOffset:  0,
		writeOffset: 0,
	}
}

// Returns the read offset
func (b *Buffer) ReadOffset() int {
	return b.readOffset
}

// Returns the write offset
func (b *Buffer) WriteOffset() int {
	return b.writeOffset
}
