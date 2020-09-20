package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *ForgeActionPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *ForgeActionPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *ForgeActionPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *ForgeActionPacket) SetSend(value bool) {
	packet.Send = value
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
