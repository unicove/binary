package bytes

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
		cap:         cap(slice),
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

// Reads the content until either EOF is reached or the maximum capacity of the provided slice
// gets fully used.
func (b *Buffer) Read(buf []byte) error {
	n := b.cap - b.readOffset
	if n < 1 {
		return EOF_ERROR
	}

	l := min(n, cap(buf))
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
