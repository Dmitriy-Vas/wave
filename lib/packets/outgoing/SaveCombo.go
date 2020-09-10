package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib"
)

// GetID returns packet ID.
func (d *SaveComboPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *SaveComboPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *SaveComboPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *SaveComboPacket) SetSend(value bool) {
	d.Send = value
}

type SaveComboPacket struct {
	ID         int64
	Send       bool
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
