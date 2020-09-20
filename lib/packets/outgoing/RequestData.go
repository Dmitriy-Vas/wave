package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *RequestDataPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *RequestDataPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *RequestDataPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *RequestDataPacket) SetSend(value bool) {
	packet.Send = value
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
