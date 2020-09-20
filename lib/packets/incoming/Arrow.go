package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib"
)

// GetID returns packet ID.
func (packet *ArrowPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *ArrowPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *ArrowPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *ArrowPacket) SetSend(value bool) {
	packet.Send = value
}

type ArrowPacket struct {
	ID     int64
	Send   bool
	Arrows []lib.ArrowRec
}

func (packet *ArrowPacket) Read(b buffer.PacketBuffer) {
	packet.Arrows = make([]lib.ArrowRec, 4)
	for i := range packet.Arrows {
		packet.Arrows[i] = lib.ArrowRec{
			Num:   b.ReadInt(b.Bytes(), b.Index()),
			Value: b.ReadInt(b.Bytes(), b.Index()),
		}
	}
}

func (packet *ArrowPacket) Write(b buffer.PacketBuffer) {
	for _, arrow := range packet.Arrows {
		b.WriteInt(b.Bytes(), arrow.Num, b.Index())
		b.WriteInt(b.Bytes(), arrow.Value, b.Index())
	}
}
