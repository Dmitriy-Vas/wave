package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *PartyRequestPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *PartyRequestPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *PartyRequestPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *PartyRequestPacket) SetSend(value bool) {
	d.Send = value
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
