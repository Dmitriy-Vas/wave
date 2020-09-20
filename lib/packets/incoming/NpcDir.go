package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib/objects"
)

// GetID returns packet ID.
func (packet *NpcDirPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *NpcDirPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *NpcDirPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *NpcDirPacket) SetSend(value bool) {
	packet.Send = value
}

type NpcDirPacket struct {
	ID         int64
	Send       bool
	NpcNum     int64
	Dir        byte
	IsPosition bool
	Position   objects.Vector2
}

func (packet *NpcDirPacket) Read(b buffer.PacketBuffer) {
	packet.NpcNum = b.ReadLong(b.Bytes(), b.Index())
	packet.Dir = b.ReadByte(b.Bytes(), b.Index())
	if packet.IsPosition = b.ReadBool(b.Bytes(), b.Index()); packet.IsPosition {
		packet.Position = objects.Vector2{
			X: b.ReadInt(b.Bytes(), b.Index()),
			Y: b.ReadInt(b.Bytes(), b.Index()),
		}
	}
}

func (packet *NpcDirPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.NpcNum, b.Index())
	b.WriteByte(b.Bytes(), packet.Dir, b.Index())
	b.WriteBool(b.Bytes(), packet.IsPosition, b.Index())
	if packet.IsPosition {
		b.WriteInt(b.Bytes(), packet.Position.X, b.Index())
		b.WriteInt(b.Bytes(), packet.Position.Y, b.Index())
	}
}
