package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *UpdateCardPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *UpdateCardPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *UpdateCardPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *UpdateCardPacket) SetSend(value bool) {
	d.Send = value
}

type UpdateCardPacket struct {
	ID   int64
	Send bool
	Num  int32
}

func (packet *UpdateCardPacket) Read(b buffer.PacketBuffer) {
	packet.Num = b.ReadInt(b.Bytes(), b.Index())
	// TODO card data
}

func (packet *UpdateCardPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Num, b.Index())
	// TODO card data
}
