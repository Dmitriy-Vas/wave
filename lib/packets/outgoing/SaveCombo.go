package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib"
)

type SaveComboPacket struct {
	*wave.DefaultPacket
	Variable1  int32
	ComboCache []lib.ComboCacheDataRec
}

func (packet *SaveComboPacket) Read(b buffer.PacketBuffer) {
	packet.ComboCache = make([]lib.ComboCacheDataRec, 3) // TODO int to const
	for i, _ := range packet.ComboCache {
		packet.ComboCache[i] = lib.ComboCacheDataRec{
			Num: b.ReadInt(b.Bytes(), b.Index()),
		}
	}
}

func (packet *SaveComboPacket) Write(b buffer.PacketBuffer) {
	for _, cache := range packet.ComboCache {
		b.WriteInt(b.Bytes(), cache.Num, b.Index())
	}
}
