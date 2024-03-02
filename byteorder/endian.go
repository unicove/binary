package byteorder

// Endian refers to the byterorder in which a sequence of bytes is stored or transmitted over the network.
type Endian = byte

const (
	// Little-endian is an order in which the little end, the least significant value in the sequence, is first
	LittleEndian Endian = 0x00
	// Big-endian is an order in which the big end, the most significant value in the sequence, is first
	BigEndian Endian = 0x01
)
