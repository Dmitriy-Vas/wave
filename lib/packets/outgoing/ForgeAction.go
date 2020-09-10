package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *ForgeActionPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *ForgeActionPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *ForgeActionPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *ForgeActionPacket) SetSend(value bool) {
	d.Send = value
}

type ForgeActionPacket struct {
	ID         int64
	Send       bool
	ForgeTimer byte
}

func (packet *ForgeActionPacket) Read(b buffer.PacketBuffer) {
	packet.ForgeTimer = b.ReadByte(b.Bytes(), b.Index())

}

func (packet *ForgeActionPacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.ForgeTimer, b.Index())
}
