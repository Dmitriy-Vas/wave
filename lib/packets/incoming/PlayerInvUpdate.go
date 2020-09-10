package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *PlayerInvUpdatePacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *PlayerInvUpdatePacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *PlayerInvUpdatePacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *PlayerInvUpdatePacket) SetSend(value bool) {
	d.Send = value
}

type PlayerInvUpdatePacket struct {
	ID        int64
	Send      bool
	Variable0 int64
	Variable1 byte
	Variable2 int32
	Variable3 int64
	Variable4 byte
	Variable5 byte
	Variable6 bool
}

func (packet *PlayerInvUpdatePacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable1 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable2 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable3 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable4 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable5 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable6 = b.ReadBool(b.Bytes(), b.Index())
}

func (packet *PlayerInvUpdatePacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable1, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable2, b.Index())
	b.WriteLong(b.Bytes(), packet.Variable3, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable4, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable5, b.Index())
	b.WriteBool(b.Bytes(), packet.Variable6, b.Index())
}
