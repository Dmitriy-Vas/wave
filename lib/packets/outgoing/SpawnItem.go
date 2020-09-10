package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *SpawnItemPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *SpawnItemPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *SpawnItemPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *SpawnItemPacket) SetSend(value bool) {
	d.Send = value
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
