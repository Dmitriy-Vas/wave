package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *SelectMinigamesPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *SelectMinigamesPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *SelectMinigamesPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *SelectMinigamesPacket) SetSend(value bool) {
	d.Send = value
}

type SelectMinigamesPacket struct {
	ID    int64
	Send  bool
	Index byte
}

func (packet *SelectMinigamesPacket) Read(b buffer.PacketBuffer) {
	packet.Index = b.ReadByte(b.Bytes(), b.Index())

}

func (packet *SelectMinigamesPacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.Index, b.Index())
}
