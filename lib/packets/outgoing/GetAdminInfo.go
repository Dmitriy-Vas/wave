package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *GetAdminInfoPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *GetAdminInfoPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *GetAdminInfoPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *GetAdminInfoPacket) SetSend(value bool) {
	packet.Send = value
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
