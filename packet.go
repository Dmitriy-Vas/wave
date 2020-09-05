package wave

import (
	. "github.com/Dmitriy-Vas/wave/buffer"
	"reflect"
)

type Packet interface {
	GetID() int64
	SetID(id int64)
	SetSend(value bool)
	GetSend() bool
	Read(b PacketBuffer)
	Write(b PacketBuffer)
}

// DefaultPacket partially implement Packet interface.
// Can be used in actual packet for shortcut.
type DefaultPacket struct {
	ID   int64
	Send bool
}

// InitPacket creates new DefaultPacket and asserts to the Packet.
func InitPacket(packet Packet) {
	dp := &DefaultPacket{Send: true}
	value := reflect.ValueOf(dp)
	reflect.ValueOf(packet).Elem().FieldByName("DefaultPacket").Set(value)
}

// GetID returns packet ID.
func (d *DefaultPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *DefaultPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *DefaultPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *DefaultPacket) SetSend(value bool) {
	d.Send = value
}
