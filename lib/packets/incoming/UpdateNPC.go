package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *UpdateNPCPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *UpdateNPCPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *UpdateNPCPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *UpdateNPCPacket) SetSend(value bool) {
	d.Send = value
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
