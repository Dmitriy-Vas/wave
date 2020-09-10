package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *RequestDataPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *RequestDataPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *RequestDataPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *RequestDataPacket) SetSend(value bool) {
	d.Send = value
}

type RequestDataPacket struct {
	ID       int64
	Send     bool
	DataType int32
}

func (packet *RequestDataPacket) Read(b buffer.PacketBuffer) {
	packet.DataType = b.ReadInt(b.Bytes(), b.Index())

}

func (packet *RequestDataPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.DataType, b.Index())
}
