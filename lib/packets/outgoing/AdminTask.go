package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *AdminTaskPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *AdminTaskPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *AdminTaskPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *AdminTaskPacket) SetSend(value bool) {
	d.Send = value
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
