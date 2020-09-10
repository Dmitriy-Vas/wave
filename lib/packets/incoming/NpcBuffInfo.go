package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *NpcBuffInfoPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *NpcBuffInfoPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *NpcBuffInfoPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *NpcBuffInfoPacket) SetSend(value bool) {
	d.Send = value
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
