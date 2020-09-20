package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *SetAccessPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *SetAccessPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *SetAccessPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *SetAccessPacket) SetSend(value bool) {
	packet.Send = value
}

type SetAccessPacket struct {
	ID     int64
	Send   bool
	Name   string
	Access int64
}

func (packet *SetAccessPacket) Read(b buffer.PacketBuffer) {
	packet.Name = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Access = b.ReadLong(b.Bytes(), b.Index())

}

func (packet *SetAccessPacket) Write(b buffer.PacketBuffer) {
	b.WriteString(b.Bytes(), packet.Name, b.Index())
	b.WriteLong(b.Bytes(), packet.Access, b.Index())
}
