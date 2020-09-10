package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *UpdatePlayerHonorPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *UpdatePlayerHonorPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *UpdatePlayerHonorPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *UpdatePlayerHonorPacket) SetSend(value bool) {
	d.Send = value
}

type UpdatePlayerHonorPacket struct {
	ID     int64
	Send   bool
	Name   string
	Target int32
}

func (packet *UpdatePlayerHonorPacket) Read(b buffer.PacketBuffer) {
	packet.Name = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Target = b.ReadInt(b.Bytes(), b.Index())

}

func (packet *UpdatePlayerHonorPacket) Write(b buffer.PacketBuffer) {
	b.WriteString(b.Bytes(), packet.Name, b.Index())
	b.WriteInt(b.Bytes(), packet.Target, b.Index())
}
