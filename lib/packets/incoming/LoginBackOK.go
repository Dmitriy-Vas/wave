package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// Empty packet
// GetID returns packet ID.
func (d *LoginBackOKPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *LoginBackOKPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *LoginBackOKPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *LoginBackOKPacket) SetSend(value bool) {
	d.Send = value
}

type LoginBackOKPacket struct {
	ID   int64
	Send bool
}

func (packet *LoginBackOKPacket) Read(b buffer.PacketBuffer) {
}

func (packet *LoginBackOKPacket) Write(b buffer.PacketBuffer) {
}
