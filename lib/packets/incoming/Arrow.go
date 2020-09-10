package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib"
)

// GetID returns packet ID.
func (d *ArrowPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *ArrowPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *ArrowPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *ArrowPacket) SetSend(value bool) {
	d.Send = value
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
