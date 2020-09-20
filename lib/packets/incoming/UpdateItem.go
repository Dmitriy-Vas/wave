package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *UpdateItemPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *UpdateItemPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *UpdateItemPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *UpdateItemPacket) SetSend(value bool) {
	packet.Send = value
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
