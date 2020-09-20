package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *UpdateNPCPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *UpdateNPCPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *UpdateNPCPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *UpdateNPCPacket) SetSend(value bool) {
	packet.Send = value
}

type UpdateNPCPacket struct {
	ID   int64
	Send bool
	Num  int32
}

func (packet *UpdateNPCPacket) Read(b buffer.PacketBuffer) {
	packet.Num = b.ReadInt(b.Bytes(), b.Index())
	// TODO NPC Data
}

func (packet *UpdateNPCPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Num, b.Index())
	// TODO NPC Data
}
