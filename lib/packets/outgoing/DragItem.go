package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *DragItemPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *DragItemPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *DragItemPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *DragItemPacket) SetSend(value bool) {
	packet.Send = value
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
