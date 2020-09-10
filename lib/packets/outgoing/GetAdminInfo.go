package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *GetAdminInfoPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *GetAdminInfoPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *GetAdminInfoPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *GetAdminInfoPacket) SetSend(value bool) {
	d.Send = value
}

type GetAdminInfoPacket struct {
	ID   int64
	Send bool
	Info int32
}

func (packet *GetAdminInfoPacket) Read(b buffer.PacketBuffer) {
	packet.Info = b.ReadInt(b.Bytes(), b.Index())

}

func (packet *GetAdminInfoPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Info, b.Index())
}
