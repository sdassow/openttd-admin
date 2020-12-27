package admin

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
)

type Packet struct {
	Type    PacketType
	Payload []byte
}

const HeaderSize int = 3

func ReadPacket(conn io.Reader) (*Packet, error) {
	// header first
	hdr := make([]byte, HeaderSize)
	_, err := conn.Read(hdr)
	if err != nil {
		return nil, err
	}

	t := PacketType(hdr[2])
	if !t.IsAPacketType() {
		return nil, fmt.Errorf("invalid packet type: %+v", t)
	}

	l := int(binary.LittleEndian.Uint16(hdr)) - HeaderSize
	p := &Packet{
		Type:    t,
		Payload: make([]byte, l),
	}

	// payload next
	_, err = conn.Read(p.Payload)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func NewPacket(t PacketType) *Packet {
	return &Packet{
		Type:    t,
		Payload: make([]byte, 0),
	}
}

func (p *Packet) Append(es ...interface{}) error {
	for _, e := range es {
		switch v := e.(type) {
		case string:
			p.Payload = append(p.Payload, []byte(v)...)
			p.Payload = append(p.Payload, byte(0))
		case uint8:
			p.Payload = append(p.Payload, byte(v))
		case uint16:
			b := make([]byte, 2)
			binary.LittleEndian.PutUint16(b, v)
			p.Payload = append(p.Payload, b...)
		case uint32:
			b := make([]byte, 4)
			binary.LittleEndian.PutUint32(b, v)
			p.Payload = append(p.Payload, b...)
		default:
			return fmt.Errorf("Unable to handle type: %+v", v)
		}
	}
	return nil
}

func (p *Packet) Marshal() []byte {
	b := make([]byte, len(p.Payload)+HeaderSize)
	binary.LittleEndian.PutUint16(b[0:], uint16(len(b)))
	b[2] = byte(p.Type)
	copy(b[HeaderSize:], p.Payload)

	return b
}

func NewAdminJoin(password, name, version string) *Packet {
	pkt := NewPacket(AdminPacketAdminJoin)
	pkt.Append(password)
	pkt.Append(name)
	pkt.Append(version)

	return pkt
}

func NewAdminQuit() *Packet {
	pkt := NewPacket(AdminPacketAdminQuit)

	return pkt
}

func NewAdminUpdateFrequency(t AdminUpdateType, freq AdminUpdateFrequency) *Packet {
	pkt := NewPacket(AdminPacketAdminUpdateFrequency)
	pkt.Append(uint16(t))
	pkt.Append(uint16(freq))

	return pkt
}

var AllCompanies uint32 = math.MaxUint32

func NewAdminPoll(t AdminUpdateType, n uint32) *Packet {
	pkt := NewPacket(AdminPacketAdminPoll)
	pkt.Append(uint8(t))
	//pkt.Append(uint32(math.MaxUint32))
	pkt.Append(uint32(n))

	return pkt
}

func NewAdminChat(a NetworkAction, destType DestType, dest uint32, msg string) *Packet {
	pkt := NewPacket(AdminPacketAdminChat)
	pkt.Append(uint8(a))
	pkt.Append(uint8(destType))
	pkt.Append(uint32(dest))
	pkt.Append(msg)

	return pkt
}

func NewAdminRcon(cmd string) *Packet {
	pkt := NewPacket(AdminPacketAdminRcon)
	pkt.Append(cmd)

	return pkt
}

func NewAdminGamescript(json string) *Packet {
	pkt := NewPacket(AdminPacketAdminGamescript)
	pkt.Append(json)

	return pkt
}

func NewAdminPing(n uint32) *Packet {
	pkt := NewPacket(AdminPacketAdminPing)
	pkt.Append(uint32(n))

	return pkt
}
