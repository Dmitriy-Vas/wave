package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *WarpMeToPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *WarpMeToPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *WarpMeToPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *WarpMeToPacket) SetSend(value bool) {
	packet.Send = value
}

type WarpMeToPacket struct {
	ID   int64
	Send bool
	Name string
}

func (packet *WarpMeToPacket) Read(b buffer.PacketBuffer) {
	packet.Name = b.ReadString(b.Bytes(), b.Index(), 0)

}

func (packet *WarpMeToPacket) Write(b buffer.PacketBuffer) {
	b.WriteString(b.Bytes(), packet.Name, b.Index())
}
