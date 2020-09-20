package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *OpenBackgroundPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *OpenBackgroundPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *OpenBackgroundPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *OpenBackgroundPacket) SetSend(value bool) {
	packet.Send = value
}

type OpenBackgroundPacket struct {
	ID         int64
	Send       bool
	Background int32
}

func (packet *OpenBackgroundPacket) Read(b buffer.PacketBuffer) {
	packet.Background = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *OpenBackgroundPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Background, b.Index())
}
