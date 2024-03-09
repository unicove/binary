package buffer

import (
	"bytes"
	"net"

	"github.com/gamevidea/binary/byteorder"
)

// ipVersion is used to represent two different versions of the Internet Protocol (IP)
type ipVersion = uint8

const (
	// ipv4 is the fourth version of the Internet Protocol, primarily used for identifying and addressing devices on a network
	// with a 32-bit address space.
	ipv4 ipVersion = net.IPv4len

	// ipv6, the sixth version of the Internet Protocol, employs a 128-bit address space and is designed to succeed IPv4, providing
	// a larger address pool to accommodate the growing number of devices connected to the internet.
	ipv6 ipVersion = net.IPv6len
)

// Reads a UDP Socket Address from the buffer and returns it
func (b *Buffer) ReadAddr(v *net.UDPAddr) error {
	ver, err := b.ReadUint8()
	if err != nil {
		return err
	}

	switch ver {
	case ipv4:
		ipBytes := make([]byte, 4)
		if err := b.Read(ipBytes); err != nil {
			return err
		}

		port, err := b.ReadUint16(byteorder.BigEndian)
		if err != nil {
			return err
		}

		v.IP = net.IPv4((-ipBytes[0]-1)&0xff, (-ipBytes[1]-1)&0xff, (-ipBytes[2]-1)&0xff, (-ipBytes[3]-1)&0xff)
		v.Port = int(port)
	case ipv6:
		if err := b.Shift(2); err != nil {
			return err
		}

		port, err := b.ReadUint16(byteorder.LittleEndian)
		if err != nil {
			return err
		}

		if err := b.Shift(4); err != nil {
			return err
		}

		v.IP = make([]byte, 16)
		if err := b.Read(v.IP); err != nil {
			return err
		}

		if err := b.Shift(4); err != nil {
			return err
		}

		v.Port = int(port)
	default:
		return CPI_ERROR
	}

	return nil
}

// Writes a UDP Socket Address to the buffer.
func (b *Buffer) WriteAddr(v *net.UDPAddr) error {
	switch len(v.IP) {
	case net.IPv4len:
		if err := b.WriteUint8(ipv4); err != nil {
			return err
		}

		ipBytes := v.IP.To4()
		if err := b.WriteUint8(^ipBytes[0]); err != nil {
			return err
		}
		if err := b.WriteUint8(^ipBytes[1]); err != nil {
			return err
		}
		if err := b.WriteUint8(^ipBytes[2]); err != nil {
			return err
		}
		if err := b.WriteUint8(^ipBytes[3]); err != nil {
			return err
		}

		if err := b.WriteUint16(uint16(v.Port), byteorder.BigEndian); err != nil {
			return err
		}
	case net.IPv6len:
		if err := b.WriteUint8(ipv6); err != nil {
			return err
		}

		if err := b.WriteInt16(23, byteorder.LittleEndian); err != nil {
			return err
		}

		if err := b.WriteUint16(uint16(v.Port), byteorder.BigEndian); err != nil {
			return err
		}

		if err := b.WriteInt32(0, byteorder.BigEndian); err != nil {
			return err
		}

		if err := b.Write(v.IP.To16()); err != nil {
			return err
		}

		if err := b.WriteInt32(0, byteorder.BigEndian); err != nil {
			return err
		}
	}

	return nil
}

// magic is unconnected message sequence which is found in every unconnected message sent in raknet
var magic = [16]byte{0x00, 0xff, 0xff, 0x00, 0xfe, 0xfe, 0xfe, 0xfe, 0xfd, 0xfd, 0xfd, 0xfd, 0x12, 0x34, 0x56, 0x78}

// Reads unconnected message sequence from the buffer and returns an error if the operation was
// unsuccessful.
func (b *Buffer) ReadMagic() error {
	if b.cap-b.offset < 16 {
		return EOF_ERROR
	}

	slice := b.slice[b.offset : b.offset+16]
	b.offset += 16

	if !bytes.Equal(slice, magic[:]) {
		return CPM_ERROR
	}

	return nil
}

// Writes the unconnected message sequence to the underlying buffer and returns an error if the operation
// was unsuccessful.
func (b *Buffer) WriteMagic() error {
	if b.cap-b.offset < 16 {
		return EOF_ERROR
	}

	slice := b.slice[b.offset : b.offset+16]
	b.offset += 16

	copy(slice, magic[:])
	return nil
}

// Regular RakNet uses 10 by default. MCPE uses 20. Configure this as appropriate
const SYSTEM_ADDRESSES_COUNT = 20

// HACK: This is the number of bytes left at which we should stop reading system addresses. This is useful
// for scenarios where a RakNet server sends less number of system addresses than the configured one. Such as a
// MCPE servers sends 20 (modern raknet servers) while traditional servers may send 10.
//
// NOTE: Make this dynamic in future in case microjang decides to alter packet fields in either of packets having,
// system addresses.
const readDeadline = 16

// Reads system addresses from the buffer into the provided slice and returns an error if the
// operation was unsuccessful.
func (b *Buffer) ReadSystemAddresses() error {
	var addr = net.UDPAddr{}

	for i := 0; i < SYSTEM_ADDRESSES_COUNT; i++ {
		if err := b.ReadAddr(&addr); err != nil {
			return err
		}

		if b.cap-b.offset == readDeadline {
			break
		}
	}

	return nil
}

// Writes system addresses from the provided slice in the underlying buffer and returns an error if the
// operation was unsuccessful.
func (b *Buffer) WriteSystemAddresses() error {
	addr := net.UDPAddr{
		IP:   net.ParseIP("255.255.255.255"),
		Port: 19132,
	}

	for i := 0; i < SYSTEM_ADDRESSES_COUNT; i++ {
		if err := b.WriteAddr(&addr); err != nil {
			return err
		}
	}

	return nil
}

// Reads the raknet pong data from the buffer into the provided slice and returns an error
// if the operation failed
func (b *Buffer) ReadPongData(buf []byte) error {
	l, err := b.ReadInt16(byteorder.BigEndian)
	if err != nil {
		return err
	}

	len := int(l)

	if b.cap-b.offset < len {
		return EOF_ERROR
	}

	copy(buf[:len], b.slice[b.offset:b.offset+len])
	b.offset += len

	return nil
}

// Writes the raknet pong data from the provided slice into the underlying buffer and returns
// an error if the operation failed
func (b *Buffer) WritePongData(buf []byte) error {
	len := len(buf)

	if err := b.WriteInt16(int16(len), byteorder.BigEndian); err != nil {
		return err
	}

	if b.cap-b.offset < len {
		return EOF_ERROR
	}

	copy(b.slice[b.offset:b.offset+len], buf[:len])
	b.offset += len

	return nil
}
