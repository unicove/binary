package buffer

// Represents a fixed size buffer with zero additional memory allocations. It provides fastest methods to read
// and write various datatypes that are serialized to and from a minecraft network wire.
type Buffer struct {
	slice  []byte
	cap    int
	len    int
	offset int
}

// Creates and returns a new Buffer by allocating a new slice of the provided length
func New(len int) *Buffer {
	return &Buffer{
		slice:  make([]byte, len),
		cap:    len,
		len:    len,
		offset: 0,
	}
}

// Creates a new buffer from the provided slice
func From(slice []byte) *Buffer {
	return &Buffer{
		slice:  slice,
		cap:    len(slice),
		len:    len(slice),
		offset: 0,
	}
}

// Returns the capacity of the buffer
func (b *Buffer) Capacity() int {
	return b.cap
}

func (b *Buffer) Length() int {
	return b.len
}

// Returns the buffer's internal cursor offset value
func (b *Buffer) Offset() int {
	return b.offset
}

// Sets the buffer's internal cursor offset to the one provided
func (b *Buffer) SetOffset(offset int) {
	b.offset = offset
}

// Returns the number of bytes left from the cursor's position to the length
// of the buffer.
func (b *Buffer) Remaining() int {
	return b.len - b.offset
}

// Resizes the buffer's internal length. This is sometimes used so that we can restrict
// the buffer from reading and writing data beyond a certain position in the cursor.
func (b *Buffer) Resize(len int) {
	b.len = len
}

// Resets the buffer's internal cursor position to 0 and resets the length back to the original
// capacity.
func (b *Buffer) Reset() {
	b.offset = 0
	b.len = b.cap
}

// Returns a reference to the buffer's internal slice allocated by the buffer when it was created.
// Any changes made to the provided slice will also reflect in the buffer.
func (b *Buffer) Slice() []byte {
	return b.slice
}

// Attempts to shift the offset index by the offset passed. Returns an error if the EOF would
// reach in doing so.
func (b *Buffer) Shift(n int) error {
	if b.len-b.offset-n < 1 {
		return ErrEndOfFile
	}

	b.offset += n
	return nil
}

// Creates a slice of the buffer's internal slice from the start till the index value passed and
// returns it.
func (b *Buffer) Get(index int) []byte {
	return b.slice[:index]
}

// Returns a slice of the buffer's internal slice from the start till the offset value of the
// buffer's internal cursor.
func (b *Buffer) Bytes() []byte {
	return b.slice[:b.offset]
}

// Reads the content until either EOF is reached or the maximum capacity of the provided slice
// gets fully used.
func (b *Buffer) Read(buf []byte) error {
	n := b.len - b.offset
	if n < 1 {
		return ErrEndOfFile
	}

	l := min(n, len(buf))
	copy(buf[:l], b.slice[b.offset:b.offset+l])

	b.offset += l
	return nil
}

// Writes the contents of the provided slice until either EOF is reached or the maximum capacity of
// the underlying buffer gets fully used.
func (b *Buffer) Write(buf []byte) error {
	n := b.len - b.offset
	if n < 1 {
		return ErrEndOfFile
	}

	l := min(n, len(buf))
	copy(b.slice[b.offset:b.offset+l], buf[:l])

	b.offset += n
	return nil
}
