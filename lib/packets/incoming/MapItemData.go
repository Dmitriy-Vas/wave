package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *MapItemDataPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *MapItemDataPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *MapItemDataPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *MapItemDataPacket) SetSend(value bool) {
	d.Send = value
}

type MapItemDataPacket struct {
	ID        int64
	Send      bool
	ItemNum   int32
	PlayerNum int32
}

func (packet *MapItemDataPacket) Read(b buffer.PacketBuffer) {
	packet.ItemNum = b.ReadInt(b.Bytes(), b.Index())
	packet.PlayerNum = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *MapItemDataPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.ItemNum, b.Index())
	b.WriteInt(b.Bytes(), packet.PlayerNum, b.Index())
}
