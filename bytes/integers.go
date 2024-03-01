package bytes

import "errors"

func (b *Buffer) ReadUint8() (v uint8, err error) {
	if b.cap-b.readOffset < 0 {
		return 0, errors.New(EOF_ERROR)
	}

	v = b.slice[b.readOffset]
	b.readOffset += 1

	return
}

func (b *Buffer) WriteUint8(v uint8) (err error) {
	if b.cap-b.writeOffset < 0 {
		return errors.New(EOF_ERROR)
	}

	b.slice[b.writeOffset] = v
	b.writeOffset += 1

	return
}
