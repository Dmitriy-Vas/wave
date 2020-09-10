package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *HighIndexPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *HighIndexPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *HighIndexPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *HighIndexPacket) SetSend(value bool) {
	d.Send = value
}

type HighIndexPacket struct {
	ID        int64
	Send      bool
	HighIndex int64
}

func (packet *HighIndexPacket) Read(b buffer.PacketBuffer) {
	packet.HighIndex = b.ReadLong(b.Bytes(), b.Index())
}

func (packet *HighIndexPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.HighIndex, b.Index())
}
