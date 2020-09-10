package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *DragItemPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *DragItemPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *DragItemPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *DragItemPacket) SetSend(value bool) {
	d.Send = value
}

type DragItemPacket struct {
	ID         int64
	Send       bool
	InvSlot    byte
	Target     byte
	TargetType byte
}

func (packet *DragItemPacket) Read(b buffer.PacketBuffer) {
	packet.InvSlot = b.ReadByte(b.Bytes(), b.Index())
	packet.Target = b.ReadByte(b.Bytes(), b.Index())
	packet.TargetType = b.ReadByte(b.Bytes(), b.Index())

}

func (packet *DragItemPacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.InvSlot, b.Index())
	b.WriteByte(b.Bytes(), packet.Target, b.Index())
	b.WriteByte(b.Bytes(), packet.TargetType, b.Index())
}
