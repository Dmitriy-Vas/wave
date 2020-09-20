package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib"
)

// GetID returns packet ID.
func (packet *PlayerElementPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *PlayerElementPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *PlayerElementPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *PlayerElementPacket) SetSend(value bool) {
	packet.Send = value
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
		for i := range packet.Elements {
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
