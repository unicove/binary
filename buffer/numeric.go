package buffer

import (
	"math"

	"github.com/gamevidea/binary/byteorder"
)

// Reads an unsigned byte and returns it
func (b *Buffer) ReadUint8() (v uint8, err error) {
	if b.cap-b.offset < 1 {
		return 0, EOF_ERROR
	}

	v = b.slice[b.offset]
	b.offset += 1

	return
}

// Writes an unsigned byte
func (b *Buffer) WriteUint8(v uint8) error {
	if b.cap-b.offset < 1 {
		return EOF_ERROR
	}

	b.slice[b.offset] = v
	b.offset += 1

	return nil
}

// Reads a signed byte and returns it
func (b *Buffer) ReadInt8() (v int8, err error) {
	if b.cap-b.offset < 1 {
		return 0, EOF_ERROR
	}

	v = int8(b.slice[b.offset])
	b.offset += 1

	return
}

// Writes a signed byte
func (b *Buffer) WriteInt8(v int8) error {
	if b.cap-b.offset < 1 {
		return EOF_ERROR
	}

	b.slice[b.offset] = byte(v)
	b.offset += 1

	return nil
}

// Reads an unsigned short and returns it
func (b *Buffer) ReadUint16(e byteorder.Endian) (v uint16, err error) {
	if b.cap-b.offset < 2 {
		return 0, EOF_ERROR
	}

	switch e {
	case byteorder.LittleEndian:
		v = uint16(b.slice[b.offset]) | uint16(b.slice[b.offset+1])<<8
	case byteorder.BigEndian:
		v = uint16(b.slice[b.offset+1]) | uint16(b.slice[b.offset])<<8
	default:
		return 0, CPE_ERROR
	}

	b.offset += 2
	return
}

// Writes an unsigned short
func (b *Buffer) WriteUint16(v uint16, e byteorder.Endian) error {
	if b.cap-b.offset < 2 {
		return EOF_ERROR
	}

	switch e {
	case byteorder.LittleEndian:
		b.slice[b.offset] = byte(v)
		b.slice[b.offset+1] = byte(v >> 8)
	case byteorder.BigEndian:
		b.slice[b.offset+1] = byte(v)
		b.slice[b.offset] = byte(v >> 8)
	default:
		return CPE_ERROR
	}

	b.offset += 2
	return nil
}

// Reads a signed short and returns it
func (b *Buffer) ReadInt16(e byteorder.Endian) (v int16, err error) {
	if b.cap-b.offset < 2 {
		return 0, EOF_ERROR
	}

	switch e {
	case byteorder.LittleEndian:
		v = int16(b.slice[b.offset]) | int16(b.slice[b.offset+1])<<8
	case byteorder.BigEndian:
		v = int16(b.slice[b.offset+1]) | int16(b.slice[b.offset])<<8
	default:
		return 0, CPE_ERROR
	}

	b.offset += 2
	return
}

// Writes a signed short
func (b *Buffer) WriteInt16(v int16, e byteorder.Endian) error {
	if b.cap-b.offset < 2 {
		return EOF_ERROR
	}

	switch e {
	case byteorder.LittleEndian:
		b.slice[b.offset] = byte(v)
		b.slice[b.offset+1] = byte(v >> 8)
	case byteorder.BigEndian:
		b.slice[b.offset+1] = byte(v)
		b.slice[b.offset] = byte(v >> 8)
	default:
		return CPE_ERROR
	}

	b.offset += 2
	return nil
}

// Reads an unsigned 24-bit integer and returns it.
func (b *Buffer) ReadUint24(e byteorder.Endian) (v uint32, err error) {
	if b.cap-b.offset < 3 {
		return 0, EOF_ERROR
	}

	switch e {
	case byteorder.LittleEndian:
		v = uint32(b.slice[b.offset]) | uint32(b.slice[b.offset+1])<<8 | uint32(b.slice[b.offset+2])<<16
	case byteorder.BigEndian:
		v = uint32(b.slice[b.offset+2]) | uint32(b.slice[b.offset+1])<<8 | uint32(b.slice[b.offset])<<16
	default:
		return 0, CPE_ERROR
	}

	b.offset += 3
	return
}

// Writes an unsigned 24-bit integer
func (b *Buffer) WriteUint24(v uint32, e byteorder.Endian) error {
	if b.cap-b.offset < 3 {
		return EOF_ERROR
	}

	switch e {
	case byteorder.LittleEndian:
		b.slice[b.offset] = byte(v)
		b.slice[b.offset+1] = byte(v >> 8)
		b.slice[b.offset+2] = byte(v >> 16)
	case byteorder.BigEndian:
		b.slice[b.offset+2] = byte(v)
		b.slice[b.offset+1] = byte(v >> 8)
		b.slice[b.offset] = byte(v >> 16)
	default:
		return CPE_ERROR
	}

	b.offset += 3
	return nil
}

// Reads an unsigned 32-bit integer and returns it.
func (b *Buffer) ReadUint32(e byteorder.Endian) (v uint32, err error) {
	if b.cap-b.offset < 4 {
		return 0, EOF_ERROR
	}

	switch e {
	case byteorder.LittleEndian:
		v = uint32(b.slice[b.offset]) | uint32(b.slice[b.offset+1])<<8 |
			uint32(b.slice[b.offset+2])<<16 | uint32(b.slice[b.offset+3])<<24
	case byteorder.BigEndian:
		v = uint32(b.slice[b.offset+3]) | uint32(b.slice[b.offset+2])<<8 |
			uint32(b.slice[b.offset+1])<<16 | uint32(b.slice[b.offset])<<24
	default:
		return 0, CPE_ERROR
	}

	b.offset += 4
	return
}

// Writes an unsigned 32-bit integer.
func (b *Buffer) WriteUint32(v uint32, e byteorder.Endian) error {
	if b.cap-b.offset < 4 {
		return EOF_ERROR
	}

	switch e {
	case byteorder.LittleEndian:
		b.slice[b.offset] = byte(v)
		b.slice[b.offset+1] = byte(v >> 8)
		b.slice[b.offset+2] = byte(v >> 16)
		b.slice[b.offset+3] = byte(v >> 24)
	case byteorder.BigEndian:
		b.slice[b.offset+3] = byte(v)
		b.slice[b.offset+2] = byte(v >> 8)
		b.slice[b.offset+1] = byte(v >> 16)
		b.slice[b.offset] = byte(v >> 24)
	default:
		return CPE_ERROR
	}

	b.offset += 4
	return nil
}

// Reads a signed 32-bit integer and returns it
func (b *Buffer) ReadInt32(e byteorder.Endian) (v int32, err error) {
	if b.cap-b.offset < 4 {
		return 0, EOF_ERROR
	}

	switch e {
	case byteorder.LittleEndian:
		v = int32(b.slice[b.offset]) | int32(b.slice[b.offset+1])<<8 |
			int32(b.slice[b.offset+2])<<16 | int32(b.slice[b.offset+3])<<24
	case byteorder.BigEndian:
		v = int32(b.slice[b.offset+3]) | int32(b.slice[b.offset+2])<<8 |
			int32(b.slice[b.offset+1])<<16 | int32(b.slice[b.offset])<<24
	default:
		return 0, CPE_ERROR
	}

	b.offset += 4
	return
}

// Writes a signed 32-bit integer
func (b *Buffer) WriteInt32(v int32, e byteorder.Endian) error {
	if b.cap-b.offset < 4 {
		return EOF_ERROR
	}

	switch e {
	case byteorder.LittleEndian:
		b.slice[b.offset] = byte(v)
		b.slice[b.offset+1] = byte(v >> 8)
		b.slice[b.offset+2] = byte(v >> 16)
		b.slice[b.offset+3] = byte(v >> 24)
	case byteorder.BigEndian:
		b.slice[b.offset+3] = byte(v)
		b.slice[b.offset+2] = byte(v >> 8)
		b.slice[b.offset+1] = byte(v >> 16)
		b.slice[b.offset] = byte(v >> 24)
	default:
		return CPE_ERROR
	}

	b.offset += 4
	return nil
}

// Reads an unsigned 64-bit integer and returns it
func (b *Buffer) ReadUint64(e byteorder.Endian) (v uint64, err error) {
	if b.cap-b.offset < 8 {
		return 0, EOF_ERROR
	}

	switch e {
	case byteorder.LittleEndian:
		v = uint64(b.slice[b.offset]) | uint64(b.slice[b.offset+1])<<8 |
			uint64(b.slice[b.offset+2])<<16 | uint64(b.slice[b.offset+3])<<24 |
			uint64(b.slice[b.offset+4])<<32 | uint64(b.slice[b.offset+5])<<40 |
			uint64(b.slice[b.offset+6])<<48 | uint64(b.slice[b.offset+7])<<56
	case byteorder.BigEndian:
		v = uint64(b.slice[b.offset+7]) | uint64(b.slice[b.offset+6])<<8 |
			uint64(b.slice[b.offset+5])<<16 | uint64(b.slice[b.offset+4])<<24 |
			uint64(b.slice[b.offset+3])<<32 | uint64(b.slice[b.offset+2])<<40 |
			uint64(b.slice[b.offset+1])<<48 | uint64(b.slice[b.offset])<<56
	default:
		return 0, CPE_ERROR
	}

	b.offset += 8
	return
}

// Writes an unsigned 64-bit integer
func (b *Buffer) WriteUint64(v uint64, e byteorder.Endian) error {
	if b.cap-b.offset < 8 {
		return EOF_ERROR
	}

	switch e {
	case byteorder.LittleEndian:
		b.slice[b.offset] = byte(v)
		b.slice[b.offset+1] = byte(v >> 8)
		b.slice[b.offset+2] = byte(v >> 16)
		b.slice[b.offset+3] = byte(v >> 24)
		b.slice[b.offset+4] = byte(v >> 32)
		b.slice[b.offset+5] = byte(v >> 40)
		b.slice[b.offset+6] = byte(v >> 48)
		b.slice[b.offset+7] = byte(v >> 56)
	case byteorder.BigEndian:
		b.slice[b.offset+7] = byte(v)
		b.slice[b.offset+6] = byte(v >> 8)
		b.slice[b.offset+5] = byte(v >> 16)
		b.slice[b.offset+4] = byte(v >> 24)
		b.slice[b.offset+3] = byte(v >> 32)
		b.slice[b.offset+2] = byte(v >> 40)
		b.slice[b.offset+1] = byte(v >> 48)
		b.slice[b.offset] = byte(v >> 56)
	default:
		return CPE_ERROR
	}

	b.offset += 8
	return nil
}

// Reads a signed 64-bit integer and returns it
func (b *Buffer) ReadInt64(e byteorder.Endian) (v int64, err error) {
	if b.cap-b.offset < 8 {
		return 0, EOF_ERROR
	}

	switch e {
	case byteorder.LittleEndian:
		v = int64(b.slice[b.offset]) | int64(b.slice[b.offset+1])<<8 |
			int64(b.slice[b.offset+2])<<16 | int64(b.slice[b.offset+3])<<24 |
			int64(b.slice[b.offset+4])<<32 | int64(b.slice[b.offset+5])<<40 |
			int64(b.slice[b.offset+6])<<48 | int64(b.slice[b.offset+7])<<56
	case byteorder.BigEndian:
		v = int64(b.slice[b.offset+7]) | int64(b.slice[b.offset+6])<<8 |
			int64(b.slice[b.offset+5])<<16 | int64(b.slice[b.offset+4])<<24 |
			int64(b.slice[b.offset+3])<<32 | int64(b.slice[b.offset+2])<<40 |
			int64(b.slice[b.offset+1])<<48 | int64(b.slice[b.offset])<<56
	default:
		return 0, CPE_ERROR
	}

	b.offset += 8
	return
}

// Writes a signed 64-bit integer
func (b *Buffer) WriteInt64(v int64, e byteorder.Endian) error {
	if b.cap-b.offset < 8 {
		return EOF_ERROR
	}

	switch e {
	case byteorder.LittleEndian:
		b.slice[b.offset] = byte(v)
		b.slice[b.offset+1] = byte(v >> 8)
		b.slice[b.offset+2] = byte(v >> 16)
		b.slice[b.offset+3] = byte(v >> 24)
		b.slice[b.offset+4] = byte(v >> 32)
		b.slice[b.offset+5] = byte(v >> 40)
		b.slice[b.offset+6] = byte(v >> 48)
		b.slice[b.offset+7] = byte(v >> 56)
	case byteorder.BigEndian:
		b.slice[b.offset+7] = byte(v)
		b.slice[b.offset+6] = byte(v >> 8)
		b.slice[b.offset+5] = byte(v >> 16)
		b.slice[b.offset+4] = byte(v >> 24)
		b.slice[b.offset+3] = byte(v >> 32)
		b.slice[b.offset+2] = byte(v >> 40)
		b.slice[b.offset+1] = byte(v >> 48)
		b.slice[b.offset] = byte(v >> 56)
	default:
		return CPE_ERROR
	}

	b.offset += 8
	return nil
}

// Reads a 32-bit floating point decimal number and returns it
func (b *Buffer) ReadFloat32(e byteorder.Endian) (v float32, err error) {
	if b.cap-b.offset < 4 {
		return 0, EOF_ERROR
	}

	switch e {
	case byteorder.LittleEndian:
		bits := uint32(b.slice[b.offset]) | uint32(b.slice[b.offset+1])<<8 |
			uint32(b.slice[b.offset+2])<<16 | uint32(b.slice[b.offset+3])<<24
		v = math.Float32frombits(bits)
	case byteorder.BigEndian:
		bits := uint32(b.slice[b.offset+3]) | uint32(b.slice[b.offset+2])<<8 |
			uint32(b.slice[b.offset+1])<<16 | uint32(b.slice[b.offset])<<24
		v = math.Float32frombits(bits)
	default:
		return 0, CPE_ERROR
	}

	b.offset += 4
	return
}

// Writes a 32-bit floating point decimal number
func (b *Buffer) WriteFloat32(v float32, e byteorder.Endian) error {
	if b.cap-b.offset < 4 {
		return EOF_ERROR
	}

	bits := math.Float32bits(v)

	switch e {
	case byteorder.LittleEndian:
		b.slice[b.offset] = byte(bits)
		b.slice[b.offset+1] = byte(bits >> 8)
		b.slice[b.offset+2] = byte(bits >> 16)
		b.slice[b.offset+3] = byte(bits >> 24)
	case byteorder.BigEndian:
		b.slice[b.offset+3] = byte(bits)
		b.slice[b.offset+2] = byte(bits >> 8)
		b.slice[b.offset+1] = byte(bits >> 16)
		b.slice[b.offset] = byte(bits >> 24)
	default:
		return CPE_ERROR
	}

	b.offset += 4
	return nil
}

// Reads a 64-bit floating point decimal number and returns it
func (b *Buffer) ReadFloat64(e byteorder.Endian) (v float64, err error) {
	if b.cap-b.offset < 8 {
		return 0, EOF_ERROR
	}

	switch e {
	case byteorder.LittleEndian:
		bits := uint64(b.slice[b.offset]) | uint64(b.slice[b.offset+1])<<8 |
			uint64(b.slice[b.offset+2])<<16 | uint64(b.slice[b.offset+3])<<24 |
			uint64(b.slice[b.offset+4])<<32 | uint64(b.slice[b.offset+5])<<40 |
			uint64(b.slice[b.offset+6])<<48 | uint64(b.slice[b.offset+7])<<56
		v = math.Float64frombits(bits)
	case byteorder.BigEndian:
		bits := uint64(b.slice[b.offset+7]) | uint64(b.slice[b.offset+6])<<8 |
			uint64(b.slice[b.offset+5])<<16 | uint64(b.slice[b.offset+4])<<24 |
			uint64(b.slice[b.offset+3])<<32 | uint64(b.slice[b.offset+2])<<40 |
			uint64(b.slice[b.offset+1])<<48 | uint64(b.slice[b.offset])<<56
		v = math.Float64frombits(bits)
	default:
		return 0, CPE_ERROR
	}

	b.offset += 8
	return
}

// Writes a 64-bit floating point decimal number
func (b *Buffer) WriteFloat64(v float64, e byteorder.Endian) error {
	if b.cap-b.offset < 8 {
		return EOF_ERROR
	}

	bits := math.Float64bits(v)

	switch e {
	case byteorder.LittleEndian:
		b.slice[b.offset] = byte(bits)
		b.slice[b.offset+1] = byte(bits >> 8)
		b.slice[b.offset+2] = byte(bits >> 16)
		b.slice[b.offset+3] = byte(bits >> 24)
		b.slice[b.offset+4] = byte(bits >> 32)
		b.slice[b.offset+5] = byte(bits >> 40)
		b.slice[b.offset+6] = byte(bits >> 48)
		b.slice[b.offset+7] = byte(bits >> 56)
	case byteorder.BigEndian:
		b.slice[b.offset+7] = byte(bits)
		b.slice[b.offset+6] = byte(bits >> 8)
		b.slice[b.offset+5] = byte(bits >> 16)
		b.slice[b.offset+4] = byte(bits >> 24)
		b.slice[b.offset+3] = byte(bits >> 32)
		b.slice[b.offset+2] = byte(bits >> 40)
		b.slice[b.offset+1] = byte(bits >> 48)
		b.slice[b.offset] = byte(bits >> 56)
	default:
		return CPE_ERROR
	}

	b.offset += 8
	return nil
}
