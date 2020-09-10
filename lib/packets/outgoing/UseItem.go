package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *UseItemPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *UseItemPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *UseItemPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *UseItemPacket) SetSend(value bool) {
	d.Send = value
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
