package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib/objects"
)

// GetID returns packet ID.
func (d *SpawnItemPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *SpawnItemPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *SpawnItemPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *SpawnItemPacket) SetSend(value bool) {
	d.Send = value
}

type SpawnItemPacket struct {
	ID         int64
	Send       bool
	HighIndex  int32
	MapItemNum byte
	ItemNum    int32
	Index      int32
	Value      int32
	Pos        objects.Vector2
}

func (packet *SpawnItemPacket) Read(b buffer.PacketBuffer) {
	packet.HighIndex = b.ReadInt(b.Bytes(), b.Index())
	packet.MapItemNum = b.ReadByte(b.Bytes(), b.Index())
	packet.ItemNum = b.ReadInt(b.Bytes(), b.Index())
	if flag := packet.MapItemNum == 0 || packet.MapItemNum > 255; !flag { // TODO rewrite
		if packet.ItemNum > 0 {
			packet.Index = b.ReadInt(b.Bytes(), b.Index())
			packet.Value = b.ReadInt(b.Bytes(), b.Index())
			packet.Pos = objects.Vector2{
				X: b.ReadInt(b.Bytes(), b.Index()),
				Y: b.ReadInt(b.Bytes(), b.Index()),
			}
		}
	}
}

func (packet *SpawnItemPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.HighIndex, b.Index())
	b.WriteByte(b.Bytes(), packet.MapItemNum, b.Index())
	b.WriteInt(b.Bytes(), packet.ItemNum, b.Index())
	if flag := packet.MapItemNum == 0 || packet.MapItemNum > 255; !flag { // TODO rewrite
		if packet.ItemNum > 0 {
			b.WriteInt(b.Bytes(), packet.Index, b.Index())
			b.WriteInt(b.Bytes(), packet.Value, b.Index())
			b.WriteInt(b.Bytes(), packet.Pos.X, b.Index())
			b.WriteInt(b.Bytes(), packet.Pos.Y, b.Index())
		}
	}
}
