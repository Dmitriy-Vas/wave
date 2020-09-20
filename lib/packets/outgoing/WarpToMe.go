package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *WarpToMePacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *WarpToMePacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *WarpToMePacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *WarpToMePacket) SetSend(value bool) {
	packet.Send = value
}

type WarpToMePacket struct {
	ID   int64
	Send bool
	Name string
}

func (packet *WarpToMePacket) Read(b buffer.PacketBuffer) {
	packet.Name = b.ReadString(b.Bytes(), b.Index(), 0)

}

func (packet *WarpToMePacket) Write(b buffer.PacketBuffer) {
	b.WriteString(b.Bytes(), packet.Name, b.Index())
}
