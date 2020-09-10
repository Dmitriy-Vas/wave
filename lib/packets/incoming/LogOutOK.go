package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// Empty packet
// GetID returns packet ID.
func (d *LogOutOKPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *LogOutOKPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *LogOutOKPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *LogOutOKPacket) SetSend(value bool) {
	d.Send = value
}

type LogOutOKPacket struct {
	ID   int64
	Send bool
}

func (packet *LogOutOKPacket) Read(b buffer.PacketBuffer) {
}

func (packet *LogOutOKPacket) Write(b buffer.PacketBuffer) {
}
