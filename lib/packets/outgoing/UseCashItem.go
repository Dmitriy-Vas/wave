package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *UseCashItemPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *UseCashItemPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *UseCashItemPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *UseCashItemPacket) SetSend(value bool) {
	d.Send = value
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
