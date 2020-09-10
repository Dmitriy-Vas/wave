package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *PlayerVitalsPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *PlayerVitalsPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *PlayerVitalsPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *PlayerVitalsPacket) SetSend(value bool) {
	d.Send = value
}

type PlayerVitalsPacket struct {
	ID        int64
	Send      bool
	Variable0 int64
	Variable1 int32
	Variable2 int64
	Variable3 int64
	Variable4 int64
	Variable5 byte
	Variable6 int32
}

func (packet *PlayerVitalsPacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable1 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable2 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable3 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable4 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable5 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable6 = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *PlayerVitalsPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable1, b.Index())
	b.WriteLong(b.Bytes(), packet.Variable2, b.Index())
	b.WriteLong(b.Bytes(), packet.Variable3, b.Index())
	b.WriteLong(b.Bytes(), packet.Variable4, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable5, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable6, b.Index())
}
