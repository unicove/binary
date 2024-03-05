package buffer

// boolByte is an alias for the byte representation of booleans. True is expressed as 0x01 and
// False is expressed as 0x00 in their respective byte representation.
type boolByte = uint8

const (
	// trueByte is the byte representation for true 0x01 in bytes
	trueByte boolByte = 0x01
	// falseByte is the byte representation for false 0x00 in bytes
	falseByte boolByte = 0x00
)

// Reads a boolean from the buffer and returns it
func (b *Buffer) ReadBool(v *bool) error {
	byte, err := b.ReadUint8()
	if err != nil {
		return err
	}

	switch byte {
	case trueByte:
		*v = true
	case falseByte:
		*v = false
	default:
		return CPB_ERROR
	}

	return nil
}

// Writes the provided boolean value into the buffer
func (b *Buffer) WriteBool(v bool) error {
	switch v {
	case true:
		return b.WriteUint8(trueByte)
	case false:
		return b.WriteUint8(falseByte)
	}

	return nil
}
