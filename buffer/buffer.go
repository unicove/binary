package buffer

// Represents a fixed size buffer with zero additional memory allocations. It provides fastest methods to read
// and write various datatypes that are serialized to and from a minecraft network wire.
type Buffer struct {
	slice  []byte
	cap    int
	len    int
	offset int
}

// Creates and returns a new Buffer of provided capacity
func New(cap int) Buffer {
	return Buffer{
		slice:  make([]byte, cap),
		cap:    cap,
		len:    cap,
		offset: 0,
	}
}

// Creates a new buffer from the provided slice
func From(slice []byte) Buffer {
	return Buffer{
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

// Returns the number of bytes left to reach the buffer's internal cursor value
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

// Returns a shared reference to the buffer's internal slice.
func (b *Buffer) Slice() []byte {
	return b.slice
}

// Returns a shared reference to the buffer's internal slice returning a sequence of bytes that
// we have read till now or written till now to the buffer.
func (b *Buffer) Bytes() []byte {
	return b.slice[:b.offset]
}

// Shifts the buffer's offset by the number of bytes passed. Returns an error if the operation
// failed.
func (b *Buffer) Shift(n int) error {
	if b.len-b.offset-n < 1 {
		return ErrEndOfFile
	}

	b.offset += n
	return nil
}

// Returns a shared reference to the buffer's internal slice containing the number of bytes passed.
func (b *Buffer) Get(bytes int) ([]byte, error) {
	n := b.len - b.offset
	if n < 1 {
		return nil, ErrEndOfFile
	}

	l := min(n, bytes)
	slice := b.slice[b.offset : b.offset+l]

	b.offset += l
	return slice, nil
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
