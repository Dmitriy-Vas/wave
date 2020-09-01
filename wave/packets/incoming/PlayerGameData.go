package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/packets/data"
)

type PlayerGameDataPacket struct {
	*wave.DefaultPacket
	PlayerCoins   int32
	PlayerHotkeys map[int32]int32
	Hotbar        map[int]Hotbar
	PlayerSpells  map[int]Spell
	PlayerVars    map[int]bool
	PlayerQuests  map[int]Quest
}

type Hotbar struct { // TODO move to models
	Slot int32
	Type byte
}

type Spell struct { // TODO move to models
	Spell  int32
	Uses   int32
	Master bool
}

type Quest struct { // TODO move to models
	QuestNum     int32
	Status       byte
	ActualTask   byte
	CurrentCount int32
	TaskCount    int32
}

func (packet *PlayerGameDataPacket) Read(b buffer.PacketBuffer) {
	packet.PlayerCoins = b.ReadInt(b.Bytes(), b.Index())
	for i := 0; i <= 59; i++ { // TODO move int to constants
		if num := b.ReadInt(b.Bytes(), b.Index()); num > 0 {
			packet.PlayerHotkeys[int32(i)] = num
		}
	}
	packet.Hotbar = make(map[int]Hotbar)
	for i := 1; i <= 12; i++ { // TODO move int to constants
		packet.Hotbar[i] = Hotbar{
			Slot: b.ReadInt(b.Bytes(), b.Index()),
			Type: b.ReadByte(b.Bytes(), b.Index()),
		}
	}
	for i := 1; i <= 49; i++ { // TODO move int to constants
		// PlayerInvData
		// TODO finish reading player inventory
		id := b.ReadInt(b.Bytes(), b.Index())
		b.ReadLong(b.Bytes(), b.Index())
		b.ReadByte(b.Bytes(), b.Index())
		if data.GetItemInt(id, data.Enhancement) > 0 {
			for f := 0; f <= 5; f++ { // TODO move int to constants
				b.ReadByte(b.Bytes(), b.Index())
			}
		}
	}
	for i := 1; i <= 240; i++ { // TODO move int to constants
		// PlayerCashInvItem
		b.ReadVector2(b.Bytes(), b.Index())
	}
	packet.PlayerSpells = make(map[int]Spell)
	for i := 1; i < 42; i++ { // TODO move int to constants
		packet.PlayerSpells[i] = Spell{
			Spell:  b.ReadInt(b.Bytes(), b.Index()),
			Uses:   b.ReadInt(b.Bytes(), b.Index()),
			Master: b.ReadBool(b.Bytes(), b.Index()),
		}
	}
	packet.PlayerVars = make(map[int]bool)
	for i := 1; i <= 100; i++ { // TODO move int to constants
		packet.PlayerVars[i] = b.ReadBool(b.Bytes(), b.Index())
	}
	_ = b.ReadLong(b.Bytes(), b.Index()) // TODO Save TNL to global
	if quests := b.ReadInt(b.Bytes(), b.Index()); quests > 0 {
		for i := 1; int32(i) <= quests; i++ {
			packet.PlayerQuests[i] = Quest{
				QuestNum:     b.ReadInt(b.Bytes(), b.Index()),
				Status:       b.ReadByte(b.Bytes(), b.Index()),
				ActualTask:   b.ReadByte(b.Bytes(), b.Index()),
				CurrentCount: b.ReadInt(b.Bytes(), b.Index()),
				TaskCount:    0,
			}
			for x := 1; x <= 10; x++ { // TODO move int to constants
				// TODO Add task counter
			}
		}
	}
	b.ReadInt(b.Bytes(), b.Index())
	b.ReadByte(b.Bytes(), b.Index())
	b.ReadInt(b.Bytes(), b.Index())

	if i := b.ReadBool(b.Bytes(), b.Index()); i {
		for x := 1; x <= 255; x++ { // TODO move int to constants
			// PlayerCraft
			b.ReadInt(b.Bytes(), b.Index())
		}
	}
	for i := 1; i <= 255; i++ { // TODO move int to constants
		// PlayerCardLevel
		b.ReadInt(b.Bytes(), b.Index())
		b.ReadInt(b.Bytes(), b.Index())
	}
	for i := 0; i <= 50; i++ { // TODO move int to constants
		// PlayerEmoji
		b.ReadInt(b.Bytes(), b.Index())
	}
	for i := 1; i <= 200; i++ { // TODO move int to constants

	}
}

func (packet *PlayerGameDataPacket) Write(b buffer.PacketBuffer) {

}
