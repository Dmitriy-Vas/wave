package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib"
)

// GetID returns packet ID.
func (packet *SaveComboPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *SaveComboPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *SaveComboPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *SaveComboPacket) SetSend(value bool) {
	packet.Send = value
}

type SaveComboPacket struct {
	ID         int64
	Send       bool
	Variable1  int32
	ComboCache []lib.ComboCacheDataRec
}

func (packet *SaveComboPacket) Read(b buffer.PacketBuffer) {
	packet.ComboCache = make([]lib.ComboCacheDataRec, 3) // TODO int to const
	for i := range packet.ComboCache {
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
