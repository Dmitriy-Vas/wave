package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *PVPPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *PVPPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *PVPPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *PVPPacket) SetSend(value bool) {
	packet.Send = value
}

type PVPPacket struct {
	ID   int64
	Send bool
	Num  int64
	Num2 int64
}

func (packet *PVPPacket) Read(b buffer.PacketBuffer) {
	packet.Num = b.ReadLong(b.Bytes(), b.Index())
	packet.Num2 = b.ReadLong(b.Bytes(), b.Index())
}

func (packet *PVPPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Num, b.Index())
	b.WriteLong(b.Bytes(), packet.Num2, b.Index())
}
