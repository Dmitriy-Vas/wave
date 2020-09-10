package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *DropCalaverasPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *DropCalaverasPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *DropCalaverasPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *DropCalaverasPacket) SetSend(value bool) {
	d.Send = value
}

type DropCalaverasPacket struct {
	ID     int64
	Send   bool
	Type   byte
	Amount int64
}

func (packet *DropCalaverasPacket) Read(b buffer.PacketBuffer) {
	packet.Type = b.ReadByte(b.Bytes(), b.Index())
	packet.Amount = b.ReadLong(b.Bytes(), b.Index())

}

func (packet *DropCalaverasPacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.Type, b.Index())
	b.WriteLong(b.Bytes(), packet.Amount, b.Index())
}
