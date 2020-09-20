package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *MapItemDataPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *MapItemDataPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *MapItemDataPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *MapItemDataPacket) SetSend(value bool) {
	packet.Send = value
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
