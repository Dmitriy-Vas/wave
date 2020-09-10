package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *PVPPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *PVPPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *PVPPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *PVPPacket) SetSend(value bool) {
	d.Send = value
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
