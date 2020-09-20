package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *DropItemPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *DropItemPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *DropItemPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *DropItemPacket) SetSend(value bool) {
	packet.Send = value
}

type DropItemPacket struct {
	ID     int64
	Send   bool
	InvNum byte
	Amount int64
}

func (packet *DropItemPacket) Read(b buffer.PacketBuffer) {
	packet.InvNum = b.ReadByte(b.Bytes(), b.Index())
	packet.Amount = b.ReadLong(b.Bytes(), b.Index())

}

func (packet *DropItemPacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.InvNum, b.Index())
	b.WriteLong(b.Bytes(), packet.Amount, b.Index())
}
