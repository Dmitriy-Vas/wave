package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib"
	"github.com/Dmitriy-Vas/wave/lib/objects"
)

type PlayerGameDataPacket struct {
	*wave.DefaultPacket
}

func (packet *PlayerGameDataPacket) Read(b buffer.PacketBuffer) {
	lib.Global.PlayerCoins = b.ReadInt(b.Bytes(), b.Index())
	for i := range lib.Global.Hotkeys {
		if num := b.ReadInt(b.Bytes(), b.Index()); num > 0 {
			lib.Global.Hotkeys[i] = num
		}
	}
	for i := range lib.Global.Hotbar {
		lib.Global.Hotbar[i] = lib.HotbarRec{
			Slot: int64(b.ReadInt(b.Bytes(), b.Index())),
			Type: b.ReadByte(b.Bytes(), b.Index()),
		}
	}
	for i := range lib.Global.PlayerInventory {
		lib.PlayerInvData(b, i)
	}
	for i := range lib.Global.PlayerCashInventory {
		vector := objects.Vector2{
			X: b.ReadInt(b.Bytes(), b.Index()),
			Y: b.ReadInt(b.Bytes(), b.Index()),
		}
		lib.SetPlayerCashInvItem(i, vector.X, vector.Y)
	}
	for i := range lib.Global.PlayerSpells {
		lib.Global.PlayerSpells[i] = lib.PlayerSpellRec{
			Spell:  b.ReadInt(b.Bytes(), b.Index()),
			Uses:   b.ReadInt(b.Bytes(), b.Index()),
			Master: b.ReadBool(b.Bytes(), b.Index()),
		}
	}
	for i := range lib.Global.PlayerVars {
		lib.Global.PlayerVars[i] = b.ReadBool(b.Bytes(), b.Index())
	}
	lib.Global.TNL = b.ReadLong(b.Bytes(), b.Index())
	if quests := b.ReadInt(b.Bytes(), b.Index()); quests > 0 {
		for i := 1; int32(i) <= quests; i++ {
			lib.Global.PlayerQuests[i] = lib.PlayerQuestRec{
				QuestNum:     b.ReadInt(b.Bytes(), b.Index()),
				Status:       b.ReadByte(b.Bytes(), b.Index()),
				ActualTask:   b.ReadByte(b.Bytes(), b.Index()),
				CurrentCount: int64(b.ReadInt(b.Bytes(), b.Index())),
				TaskCount:    0,
			}
			for x := 1; x <= 10; x++ { // TODO move int to constants
				if lib.Quest[lib.Global.PlayerQuests[i].QuestNum].Task[x].Order > 0 {
					lib.Global.PlayerQuests[i].TaskCount++
				}
			}
		}
	}
	// TODO Update available quests
	b.ReadInt(b.Bytes(), b.Index())
	b.ReadByte(b.Bytes(), b.Index())
	b.ReadInt(b.Bytes(), b.Index())

	if value := b.ReadInt(b.Bytes(), b.Index()); value > 0 {
		for i := range lib.Global.PlayerCraft {
			lib.Global.PlayerCraft[i] = b.ReadInt(b.Bytes(), b.Index())
		}
	}
	for i := range lib.Global.PlayerCards {
		lib.Global.PlayerCards[i].Level = b.ReadInt(b.Bytes(), b.Index())
		lib.Global.PlayerCards[i].Exp = b.ReadInt(b.Bytes(), b.Index())
	}
	for i := 0; i <= 50; i++ {
		b.ReadInt(b.Bytes(), b.Index()) // TODO player emojis
	}
	for i := range lib.Global.PlayerAwards {
		if lib.Global.PlayerAwards[i].Level == 0 {
			lib.Global.PlayerAwards[i].Level++
		}
		if value := b.ReadInt(b.Bytes(), b.Index()); value > 0 {
			lib.Global.PlayerAwards[value].Level = b.ReadInt(b.Bytes(), b.Index())
			lib.Global.PlayerAwards[value].Count = b.ReadInt(b.Bytes(), b.Index())
			lib.Global.PlayerAwards[value].GetDate = b.ReadString(b.Bytes(), b.Index(), 0)
		} else {
			// Garbage
			// b.ReadInt(b.Bytes(), b.Index())
		}
	}
	for i := range lib.Global.PlayerCalaveras {
		lib.Global.PlayerCalaveras[i] = b.ReadLong(b.Bytes(), b.Index())
	}
	for i := range lib.Global.PlayerPin {
		lib.Global.PlayerPin[i] = b.ReadBool(b.Bytes(), b.Index())
	}
	for i := range lib.Global.PlayerProfession {
		b.ReadByte(b.Bytes(), b.Index())
		b.ReadInt(b.Bytes(), b.Index())
		b.ReadByte(b.Bytes(), b.Index())
		b.ReadLong(b.Bytes(), b.Index())
		for x := 0; x <= 9; x++ { // TODO move int to constants
			lib.Global.PlayerProfession[i].Upgrade[x] = b.ReadByte(b.Bytes(), b.Index())
		}
	}
	b.ReadString(b.Bytes(), b.Index(), 0)
	b.ReadByte(b.Bytes(), b.Index())
	b.ReadByte(b.Bytes(), b.Index())
}

func (packet *PlayerGameDataPacket) Write(b buffer.PacketBuffer) {

}
