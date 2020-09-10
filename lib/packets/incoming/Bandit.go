package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *BanditPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *BanditPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *BanditPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *BanditPacket) SetSend(value bool) {
	d.Send = value
}

type BanditPacket struct {
	ID           int64
	Send         bool
	BanditSelect int64
}

func (packet *BanditPacket) Read(b buffer.PacketBuffer) {
	packet.BanditSelect = b.ReadLong(b.Bytes(), b.Index())
}

func (packet *BanditPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.BanditSelect, b.Index())
}
