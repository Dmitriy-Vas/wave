package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib"
)

// GetID returns packet ID.
func (d *PlayerElementPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *PlayerElementPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *PlayerElementPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *PlayerElementPacket) SetSend(value bool) {
	d.Send = value
}

type PlayerElementPacket struct {
	ID           int64
	Send         bool
	Variable1    int32
	ElementCount int32
	Elements     []lib.PlayerElementRec
}

func (packet *PlayerElementPacket) Read(b buffer.PacketBuffer) {
	packet.Variable1 = b.ReadInt(b.Bytes(), b.Index())
	packet.ElementCount = b.ReadInt(b.Bytes(), b.Index())
	if packet.ElementCount > 0 {
		packet.Elements = make([]lib.PlayerElementRec, packet.ElementCount)
		for i, _ := range packet.Elements {
			packet.Elements[i] = lib.PlayerElementRec{
				Num:  b.ReadInt(b.Bytes(), b.Index()),
				Icon: b.ReadInt(b.Bytes(), b.Index()),
			}
		}
	}
}

func (packet *PlayerElementPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Variable1, b.Index())
	b.WriteInt(b.Bytes(), packet.ElementCount, b.Index())
	if packet.ElementCount > 0 {
		for _, element := range packet.Elements {
			b.WriteInt(b.Bytes(), element.Num, b.Index())
			b.WriteInt(b.Bytes(), element.Icon, b.Index())
		}
	}
}
