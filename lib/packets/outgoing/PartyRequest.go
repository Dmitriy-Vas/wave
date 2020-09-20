package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *PartyRequestPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *PartyRequestPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *PartyRequestPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *PartyRequestPacket) SetSend(value bool) {
	packet.Send = value
}

type PartyRequestPacket struct {
	ID    int64
	Send  bool
	Index int32
}

func (packet *PartyRequestPacket) Read(b buffer.PacketBuffer) {
	packet.Index = b.ReadInt(b.Bytes(), b.Index())

}

func (packet *PartyRequestPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Index, b.Index())
}
