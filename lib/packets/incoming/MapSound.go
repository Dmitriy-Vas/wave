package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *MapSoundPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *MapSoundPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *MapSoundPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *MapSoundPacket) SetSend(value bool) {
	packet.Send = value
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
