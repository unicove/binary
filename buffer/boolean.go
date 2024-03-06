package buffer

const (
	// trueByte is the byte representation for true 0x01 in bytes
	trueByte uint8 = 0x01
	// falseByte is the byte representation for false 0x00 in bytes
	falseByte uint8 = 0x00
)

// Reads a boolean from the buffer and returns it
func (b *Buffer) ReadBool() (bool, error) {
	byte, err := b.ReadUint8()
	if err != nil {
		return false, err
	}

	switch byte {
	case trueByte:
		return true, nil
	case falseByte:
		return false, nil
	default:
		return false, CPB_ERROR
	}
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
