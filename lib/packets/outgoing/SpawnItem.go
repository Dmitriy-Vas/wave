package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *SpawnItemPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *SpawnItemPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *SpawnItemPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *SpawnItemPacket) SetSend(value bool) {
	packet.Send = value
}

type SpawnItemPacket struct {
	ID        int64
	Send      bool
	TmpItem   int32
	TmpAmount int32
}

func (packet *SpawnItemPacket) Read(b buffer.PacketBuffer) {
	packet.TmpItem = b.ReadInt(b.Bytes(), b.Index())
	packet.TmpAmount = b.ReadInt(b.Bytes(), b.Index())

}

func (packet *SpawnItemPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.TmpItem, b.Index())
	b.WriteInt(b.Bytes(), packet.TmpAmount, b.Index())
}
