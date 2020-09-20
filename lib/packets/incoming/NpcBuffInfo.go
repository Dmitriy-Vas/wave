package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *NpcBuffInfoPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *NpcBuffInfoPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *NpcBuffInfoPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *NpcBuffInfoPacket) SetSend(value bool) {
	packet.Send = value
}

type NpcBuffInfoPacket struct {
	ID    int64
	Send  bool
	Index int32
	Slot  byte
}

func (packet *NpcBuffInfoPacket) Read(b buffer.PacketBuffer) {
	packet.Index = b.ReadInt(b.Bytes(), b.Index())
	packet.Slot = b.ReadByte(b.Bytes(), b.Index())
}

func (packet *NpcBuffInfoPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Index, b.Index())
	b.WriteByte(b.Bytes(), packet.Slot, b.Index())
}
