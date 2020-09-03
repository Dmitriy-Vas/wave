package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib"
)

type UpdateCraftingPacket struct {
	*wave.DefaultPacket
	Num   int64
	Craft lib.CraftingRec
}

func (packet *UpdateCraftingPacket) Read(b buffer.PacketBuffer) {
	if packet.Num = b.ReadLong(b.Bytes(), b.Index()); packet.Num >= 0 && packet.Num <= 255 {
		craft := lib.CraftingRec{
			Name:    b.ReadString(b.Bytes(), b.Index(), 0),
			Desc:    b.ReadString(b.Bytes(), b.Index(), 0),
			ItemReq: make([]int32, 3),
			ItemVal: make([]int32, 3),
		}
		for i := range craft.ItemReq {
			craft.ItemReq[i] = b.ReadInt(b.Bytes(), b.Index())
			craft.ItemVal[i] = b.ReadInt(b.Bytes(), b.Index())
		}
		craft.Reward = b.ReadInt(b.Bytes(), b.Index())
		craft.RewardVal = b.ReadInt(b.Bytes(), b.Index())
		packet.Craft = craft
	}
}

func (packet *UpdateCraftingPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Num, b.Index())
	if packet.Num >= 0 && packet.Num <= 255 {
		b.WriteString(b.Bytes(), packet.Craft.Name, b.Index())
		b.WriteString(b.Bytes(), packet.Craft.Desc, b.Index())
		for i := range packet.Craft.ItemReq {
			b.WriteInt(b.Bytes(), packet.Craft.ItemReq[i], b.Index())
			b.WriteInt(b.Bytes(), packet.Craft.ItemVal[i], b.Index())
		}
		b.WriteInt(b.Bytes(), packet.Craft.Reward, b.Index())
		b.WriteInt(b.Bytes(), packet.Craft.RewardVal, b.Index())
	}
}
