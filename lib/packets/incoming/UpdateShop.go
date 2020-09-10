package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *UpdateShopPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *UpdateShopPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *UpdateShopPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *UpdateShopPacket) SetSend(value bool) {
	d.Send = value
}

type UpdateShopPacket struct {
	ID   int64
	Send bool
	Num  int32
}

func (packet *UpdateShopPacket) Read(b buffer.PacketBuffer) {
	packet.Num = b.ReadInt(b.Bytes(), b.Index())
	// TODO Shop data
}

func (packet *UpdateShopPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Num, b.Index())
	// TODO Shop data
}
