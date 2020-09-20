package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *UseItemPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *UseItemPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *UseItemPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *UseItemPacket) SetSend(value bool) {
	packet.Send = value
}

type UseItemPacket struct {
	ID     int64
	Send   bool
	InvNum int32
}

func (packet *UseItemPacket) Read(b buffer.PacketBuffer) {
	packet.InvNum = b.ReadInt(b.Bytes(), b.Index())

}

func (packet *UseItemPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.InvNum, b.Index())
}
