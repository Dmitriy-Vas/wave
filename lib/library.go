package lib

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib/objects"
)

func PlayerInvData(b buffer.PacketBuffer, slot int) {
	Global.PlayerInventory[slot].Num = b.ReadInt(b.Bytes(), b.Index())
	Global.PlayerInventory[slot].Value = b.ReadLong(b.Bytes(), b.Index())
	Global.PlayerInventory[slot].Slot = b.ReadByte(b.Bytes(), b.Index())
	if GetItemInt(Global.PlayerInventory[slot].Num, ItemEnhancement) > 0 {
		Global.PlayerInventory[slot].Stat = make([]byte, 5)
		for i, _ := range Global.PlayerInventory[slot].Stat {
			Global.PlayerInventory[slot].Stat[i] = b.ReadByte(b.Bytes(), b.Index())
		}
	}
}

func GetItemInt(itemNum int32, property ItemPropertyType) int32 {
	if itemNum == 0 || itemNum > 1200 {
		return 0
	}
	if property < ItemLight || property > Itemexp_lost {
		return 0
	}
	return int32(Item[itemNum].Int[property])
}

func SetPlayerCashInvItem(slot int, itemNum int32, itemValue int32) {
	Global.PlayerCashInventory[slot].Num = itemNum
	Global.PlayerCashInventory[slot].Value = int64(itemValue)
}

func ReadPlayerBankData(b buffer.PacketBuffer) InvItemRec {
	item := InvItemRec{
		Num:   b.ReadInt(b.Bytes(), b.Index()),
		Value: b.ReadLong(b.Bytes(), b.Index()),
		Slot:  b.ReadByte(b.Bytes(), b.Index()),
		Stat:  make([]byte, 6),
	}
	if true { //modTypes.Item[ptr.Num].Int[13] > 0; TODO
		for i, _ := range item.Stat {
			item.Stat[i] = b.ReadByte(b.Bytes(), b.Index())
		}
	}
	return item
}

func WritePlayerBankData(b buffer.PacketBuffer, item InvItemRec) {
	b.WriteInt(b.Bytes(), item.Num, b.Index())
	b.WriteLong(b.Bytes(), item.Value, b.Index())
	b.WriteByte(b.Bytes(), item.Slot, b.Index())
	if true { //modTypes.Item[ptr.Num].Int[13] > 0; TODO
		for _, data := range item.Stat {
			b.WriteByte(b.Bytes(), data, b.Index())
		}
	}
}

func ReadCashShopData(b buffer.PacketBuffer) CashShopRec {
	item := CashShopRec{
		Name: b.ReadString(b.Bytes(), b.Index(), 0),
		Icon: b.ReadInt(b.Bytes(), b.Index()),
	}
	item.Slot = make([]TradeCashItemRec, b.ReadInt(b.Bytes(), b.Index()))
	for i, _ := range item.Slot {
		item.Slot[i].Item = b.ReadInt(b.Bytes(), b.Index())
		item.Slot[i].Value = b.ReadLong(b.Bytes(), b.Index())
		item.Slot[i].Price = b.ReadLong(b.Bytes(), b.Index())
	}
	return item
}

func WriteCashShopData(b buffer.PacketBuffer, item CashShopRec) {
	b.WriteString(b.Bytes(), item.Name, b.Index())
	b.WriteInt(b.Bytes(), item.Icon, b.Index())
	b.WriteInt(b.Bytes(), int32(len(item.Slot)), b.Index())
	for _, slot := range item.Slot {
		b.WriteInt(b.Bytes(), slot.Item, b.Index())
		b.WriteLong(b.Bytes(), slot.Value, b.Index())
		b.WriteLong(b.Bytes(), slot.Price, b.Index())
	}
}

func ReadMapData(b buffer.PacketBuffer) MapRec {
	rec := MapRec{
		Name:         b.ReadString(b.Bytes(), b.Index(), 0),
		ESName:       b.ReadString(b.Bytes(), b.Index(), 0),
		Revision:     b.ReadInt(b.Bytes(), b.Index()),
		Moral:        b.ReadInt(b.Bytes(), b.Index()),
		MaxX:         b.ReadByte(b.Bytes(), b.Index()),
		MaxY:         b.ReadByte(b.Bytes(), b.Index()),
		Region:       b.ReadInt(b.Bytes(), b.Index()),
		Image:        b.ReadInt(b.Bytes(), b.Index()),
		Up:           b.ReadInt(b.Bytes(), b.Index()),
		Down:         b.ReadInt(b.Bytes(), b.Index()),
		Left:         b.ReadInt(b.Bytes(), b.Index()),
		Right:        b.ReadInt(b.Bytes(), b.Index()),
		BootMap:      b.ReadInt(b.Bytes(), b.Index()),
		BootX:        b.ReadInt(b.Bytes(), b.Index()),
		BootY:        b.ReadInt(b.Bytes(), b.Index()),
		Fog:          b.ReadInt(b.Bytes(), b.Index()),
		FogOpacity:   b.ReadByte(b.Bytes(), b.Index()),
		FogSpeed:     b.ReadInt(b.Bytes(), b.Index()),
		tileA:        b.ReadInt(b.Bytes(), b.Index()),
		tileB:        b.ReadInt(b.Bytes(), b.Index()),
		DayNight:     b.ReadInt(b.Bytes(), b.Index()),
		Parallax:     b.ReadInt(b.Bytes(), b.Index()),
		ParallaxType: b.ReadByte(b.Bytes(), b.Index()),
		Ambience:     b.ReadString(b.Bytes(), b.Index(), 0),
		NightAlpha:   b.ReadInt(b.Bytes(), b.Index()),
		InstanceMax:  b.ReadInt(b.Bytes(), b.Index()),
		MapCondition: b.ReadInt(b.Bytes(), b.Index()),
		Puzzle: objects.Vector4{
			A: b.ReadInt(b.Bytes(), b.Index()),
			B: b.ReadInt(b.Bytes(), b.Index()),
			C: b.ReadInt(b.Bytes(), b.Index()),
			D: b.ReadInt(b.Bytes(), b.Index()),
		},
		Drill: objects.Vector4{
			A: b.ReadInt(b.Bytes(), b.Index()),
			B: b.ReadInt(b.Bytes(), b.Index()),
			C: b.ReadInt(b.Bytes(), b.Index()),
			D: b.ReadInt(b.Bytes(), b.Index()),
		},
		tileOffset: objects.Vector2{
			X: b.ReadInt(b.Bytes(), b.Index()),
			Y: b.ReadInt(b.Bytes(), b.Index()),
		},
	}
	// TODO finish map data
	return rec
}

func WriteMapData(b buffer.PacketBuffer, item MapRec) {

}
