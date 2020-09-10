package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *UpdateItemPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *UpdateItemPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *UpdateItemPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *UpdateItemPacket) SetSend(value bool) {
	d.Send = value
}

type UpdateItemPacket struct {
	ID      int64
	Send    bool
	ItemNum int32
}

func (packet *UpdateItemPacket) Read(b buffer.PacketBuffer) {
	packet.ItemNum = b.ReadInt(b.Bytes(), b.Index())
	// TODO item data
}

func (packet *UpdateItemPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.ItemNum, b.Index())
	// TODO item data
}
