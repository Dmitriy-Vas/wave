package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *DropCalaverasPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *DropCalaverasPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *DropCalaverasPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *DropCalaverasPacket) SetSend(value bool) {
	packet.Send = value
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
