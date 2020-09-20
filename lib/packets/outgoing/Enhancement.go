package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *EnhancementPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *EnhancementPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *EnhancementPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *EnhancementPacket) SetSend(value bool) {
	packet.Send = value
}

type EnhancementPacket struct {
	ID         int64
	Send       bool
	InvSlot    byte
	TargetSlot byte
}

func (packet *EnhancementPacket) Read(b buffer.PacketBuffer) {
	packet.InvSlot = b.ReadByte(b.Bytes(), b.Index())
	packet.TargetSlot = b.ReadByte(b.Bytes(), b.Index())

}

func (packet *EnhancementPacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.InvSlot, b.Index())
	b.WriteByte(b.Bytes(), packet.TargetSlot, b.Index())
}
