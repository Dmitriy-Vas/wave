package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *SetAccessPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *SetAccessPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *SetAccessPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *SetAccessPacket) SetSend(value bool) {
	d.Send = value
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
