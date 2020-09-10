package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib/objects"
)

// GetID returns packet ID.
func (d *MapNPCDataPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *MapNPCDataPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *MapNPCDataPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *MapNPCDataPacket) SetSend(value bool) {
	d.Send = value
}

type MapNPCDataPacket struct {
	ID      int64
	Send    bool
	Num     int32
	Pos     objects.Vector2
	Dir     byte
	Vital   int32
	HPSetTo int32
}

func (packet *MapNPCDataPacket) Read(b buffer.PacketBuffer) {
	packet.Num = b.ReadInt(b.Bytes(), b.Index())
	packet.Pos = objects.Vector2{
		X: b.ReadInt(b.Bytes(), b.Index()),
		Y: b.ReadInt(b.Bytes(), b.Index()),
	}
	packet.Dir = b.ReadByte(b.Bytes(), b.Index())
	packet.Vital = b.ReadInt(b.Bytes(), b.Index())
	packet.HPSetTo = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *MapNPCDataPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Num, b.Index())
	b.WriteInt(b.Bytes(), packet.Pos.X, b.Index())
	b.WriteInt(b.Bytes(), packet.Pos.Y, b.Index())
	b.WriteByte(b.Bytes(), packet.Dir, b.Index())
	b.WriteInt(b.Bytes(), packet.Vital, b.Index())
	b.WriteInt(b.Bytes(), packet.HPSetTo, b.Index())
}
