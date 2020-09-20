package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *UpdateTitlePacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *UpdateTitlePacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *UpdateTitlePacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *UpdateTitlePacket) SetSend(value bool) {
	packet.Send = value
}

type UpdateTitlePacket struct {
	ID   int64
	Send bool
	Num  int64
	Num2 int64
	Num3 int64
}

func (packet *UpdateTitlePacket) Read(b buffer.PacketBuffer) {
	packet.Num = b.ReadLong(b.Bytes(), b.Index())
	packet.Num2 = b.ReadLong(b.Bytes(), b.Index())
	packet.Num3 = b.ReadLong(b.Bytes(), b.Index())
}

func (packet *UpdateTitlePacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Num, b.Index())
	b.WriteLong(b.Bytes(), packet.Num2, b.Index())
	b.WriteLong(b.Bytes(), packet.Num3, b.Index())
}
