package bytes

import (
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
		octets := [4]byte{}
		if err := b.Read(octets[:]); err != nil {
			return err
		}

		port, err := b.ReadUint16(byteorder.BigEndian)
		if err != nil {
			return err
		}

		v.IP = octets[:]
		v.Port = int(port)
	case ipv6:
		if err := b.Next(2); err != nil {
			return err
		}

		port, err := b.ReadUint16(byteorder.LittleEndian)
		if err != nil {
			return err
		}

		if err := b.Next(4); err != nil {
			return err
		}

		octets := [16]byte{}
		if err := b.Read(octets[:]); err != nil {
			return err
		}

		if err := b.Next(4); err != nil {
			return err
		}

		v.IP = octets[:]
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

		if err := b.Write(v.IP.To4()); err != nil {
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
