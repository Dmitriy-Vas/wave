package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib"
	"github.com/Dmitriy-Vas/wave/lib/objects"
)

type SpawnMobItemPacket struct {
	*wave.DefaultPacket
	ItemIndex int32
	MapItem   lib.MapItemRec
}

func (packet *SpawnMobItemPacket) Read(b buffer.PacketBuffer) {
	if packet.ItemIndex = b.ReadInt(b.Bytes(), b.Index()); packet.ItemIndex != 0 && packet.ItemIndex <= 255 {
		item := lib.MapItemRec{
			Item: lib.InvItemRec{
				Num: b.ReadInt(b.Bytes(), b.Index()),
			},
		}
		if item.Item.Num > 0 {
			item.Item.Value = int64(b.ReadInt(b.Bytes(), b.Index()))
			item.PIndex = b.ReadInt(b.Bytes(), b.Index())
			item.Pos = objects.Vector2{
				X: b.ReadInt(b.Bytes(), b.Index()),
				Y: b.ReadInt(b.Bytes(), b.Index()),
			}
			item.X2 = int64(b.ReadInt(b.Bytes(), b.Index()))
			item.Y2 = int64(b.ReadInt(b.Bytes(), b.Index()))
		}
		packet.MapItem = item
	}
}

func (packet *SpawnMobItemPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.ItemIndex, b.Index())
	if packet.ItemIndex != 0 && packet.ItemIndex <= 255 {
		b.WriteInt(b.Bytes(), packet.MapItem.Item.Num, b.Index())
		if packet.MapItem.Item.Num > 0 {
			b.WriteInt(b.Bytes(), int32(packet.MapItem.Item.Value), b.Index())
			b.WriteInt(b.Bytes(), packet.MapItem.PIndex, b.Index())
			b.WriteInt(b.Bytes(), packet.MapItem.Pos.X, b.Index())
			b.WriteInt(b.Bytes(), packet.MapItem.Pos.Y, b.Index())
			b.WriteInt(b.Bytes(), int32(packet.MapItem.X2), b.Index())
			b.WriteInt(b.Bytes(), int32(packet.MapItem.Y2), b.Index())
		}
	}
}
