package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib/objects"
)

// GetID returns packet ID.
func (d *MapNpcVitalsPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *MapNpcVitalsPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *MapNpcVitalsPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *MapNpcVitalsPacket) SetSend(value bool) {
	d.Send = value
}

type MapNpcVitalsPacket struct {
	ID        int64
	Send      bool
	Variable0 int64
	Variable1 int32
	Variable2 bool
	Variable3 int32
	Variable4 int32
	Variable5 byte
	Variable6 objects.Vector2
	Variable7 int32
	Variable8 int32
}

func (packet *MapNpcVitalsPacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable1 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable2 = b.ReadBool(b.Bytes(), b.Index())
	packet.Variable3 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable4 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable5 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable6 = objects.Vector2{
		X: b.ReadInt(b.Bytes(), b.Index()),
		Y: b.ReadInt(b.Bytes(), b.Index()),
	}
	packet.Variable7 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable8 = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *MapNpcVitalsPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable1, b.Index())
	b.WriteBool(b.Bytes(), packet.Variable2, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable3, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable4, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable5, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable6.X, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable6.Y, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable7, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable8, b.Index())
}
