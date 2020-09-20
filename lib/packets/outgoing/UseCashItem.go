package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *UseCashItemPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *UseCashItemPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *UseCashItemPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *UseCashItemPacket) SetSend(value bool) {
	packet.Send = value
}

type UseCashItemPacket struct {
	ID     int64
	Send   bool
	InvNum byte
}

func (packet *UseCashItemPacket) Read(b buffer.PacketBuffer) {
	packet.InvNum = b.ReadByte(b.Bytes(), b.Index())

}

func (packet *UseCashItemPacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.InvNum, b.Index())
}
