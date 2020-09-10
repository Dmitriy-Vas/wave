package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *MapSoundPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *MapSoundPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *MapSoundPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *MapSoundPacket) SetSend(value bool) {
	d.Send = value
}

type MapSoundPacket struct {
	ID    int64
	Send  bool
	Num   int32
	Sound string
}

func (packet *MapSoundPacket) Read(b buffer.PacketBuffer) {
	packet.Num = b.ReadInt(b.Bytes(), b.Index())
	packet.Sound = b.ReadString(b.Bytes(), b.Index(), 0)
}

func (packet *MapSoundPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Num, b.Index())
	b.WriteString(b.Bytes(), packet.Sound, b.Index())
}
