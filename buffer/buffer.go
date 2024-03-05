package buffer

// Represents a fixed size buffer with zero additional memory allocations. It provides fast methods to read
// and write various datatypes that are serialized to and from a minecraft network wire.
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
		cap:         len(slice),
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

// Attempts to shift the read offset index by the number passed. Returns an error if the EOF
// would reach.
func (b *Buffer) Next(n int) error {
	if b.cap-b.readOffset-n < 1 {
		return EOF_ERROR
	}

	b.readOffset += n
	return nil
}

// Returns the number of bytes remaining to be read from the buffer.
func (b *Buffer) Remaining() int {
	return b.cap - b.readOffset
}

// Reads the content until either EOF is reached or the maximum capacity of the provided slice
// gets fully used.
func (b *Buffer) Read(buf []byte) error {
	n := b.cap - b.readOffset
	if n < 1 {
		return EOF_ERROR
	}

	l := min(n, len(buf))
	copy(buf[:l], b.slice[b.readOffset:b.readOffset+l])

	b.readOffset += l
	return nil
}

// Writes the contents of the provided slice until either EOF is reached or the maximum capacity of
// the underlying buffer gets fully used.
func (b *Buffer) Write(buf []byte) error {
	n := b.cap - b.writeOffset
	if n < 1 {
		return EOF_ERROR
	}

	l := min(n, len(buf))
	copy(b.slice[b.writeOffset:b.writeOffset+l], buf[:l])

	b.writeOffset += n
	return nil
}
