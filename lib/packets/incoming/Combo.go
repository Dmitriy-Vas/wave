package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib"
)

type ComboPacket struct {
	*wave.DefaultPacket
	ComboCount  int32
	PlayerCombo []lib.PlayerComboRec
	ComboCache  []lib.ComboCacheDataRec
}

func (packet *ComboPacket) Read(b buffer.PacketBuffer) {
	packet.ComboCount = b.ReadInt(b.Bytes(), b.Index())
	packet.PlayerCombo = make([]lib.PlayerComboRec, 18) // TODO int to const
	for i, _ := range packet.PlayerCombo {
		packet.PlayerCombo[i].Num = b.ReadInt(b.Bytes(), b.Index())
	}
	packet.ComboCache = make([]lib.ComboCacheDataRec, 3) // TODO int to const
	for i, _ := range packet.ComboCache {
		packet.ComboCache[i].Num = b.ReadInt(b.Bytes(), b.Index())
	}
}

func (packet *ComboPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.ComboCount, b.Index())
	for _, combo := range packet.PlayerCombo {
		b.WriteInt(b.Bytes(), combo.Num, b.Index())
	}
	for _, cache := range packet.ComboCache {
		b.WriteInt(b.Bytes(), cache.Num, b.Index())
	}
}
