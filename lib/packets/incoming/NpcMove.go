package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib/objects"
)

// GetID returns packet ID.
func (packet *NpcMovePacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *NpcMovePacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *NpcMovePacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *NpcMovePacket) SetSend(value bool) {
	packet.Send = value
}

type NpcMovePacket struct {
	ID        int64
	Send      bool
	Variable0 int64
	Variable1 int32
	Variable2 int32
	Variable3 objects.Vector2
	Variable4 byte
	Variable5 int64
	Variable6 byte
	Variable7 byte
	Variable8 objects.Vector2
	Variable9 int32
}

func (packet *NpcMovePacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable1 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable2 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable3 = objects.Vector2{
		X: b.ReadInt(b.Bytes(), b.Index()),
		Y: b.ReadInt(b.Bytes(), b.Index()),
	}
	packet.Variable4 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable5 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable6 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable7 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable8 = objects.Vector2{
		X: b.ReadInt(b.Bytes(), b.Index()),
		Y: b.ReadInt(b.Bytes(), b.Index()),
	}
	packet.Variable9 = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *NpcMovePacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable1, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable2, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable3.X, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable3.Y, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable4, b.Index())
	b.WriteLong(b.Bytes(), packet.Variable5, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable6, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable7, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable8.X, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable8.Y, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable9, b.Index())
}
