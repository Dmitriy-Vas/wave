package lib

import "github.com/Dmitriy-Vas/wave/buffer"

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
