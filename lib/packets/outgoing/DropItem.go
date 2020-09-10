package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *DropItemPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *DropItemPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *DropItemPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *DropItemPacket) SetSend(value bool) {
	d.Send = value
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
