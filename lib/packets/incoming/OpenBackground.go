package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *OpenBackgroundPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *OpenBackgroundPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *OpenBackgroundPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *OpenBackgroundPacket) SetSend(value bool) {
	d.Send = value
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
