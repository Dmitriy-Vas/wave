package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *AdminTaskPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *AdminTaskPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *AdminTaskPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *AdminTaskPacket) SetSend(value bool) {
	packet.Send = value
}

type AdminTaskPacket struct {
	ID        int64
	Send      bool
	TaskNum   byte
	Variable2 int32
	Variable3 string
}

func (packet *AdminTaskPacket) Read(b buffer.PacketBuffer) {
	packet.TaskNum = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable2 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable3 = b.ReadString(b.Bytes(), b.Index(), 0)

}

func (packet *AdminTaskPacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.TaskNum, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable2, b.Index())
	b.WriteString(b.Bytes(), packet.Variable3, b.Index())
}
