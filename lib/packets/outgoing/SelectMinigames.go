package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *SelectMinigamesPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *SelectMinigamesPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *SelectMinigamesPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *SelectMinigamesPacket) SetSend(value bool) {
	packet.Send = value
}

type SelectMinigamesPacket struct {
	ID    int64
	Send  bool
	Index byte
}

func (packet *SelectMinigamesPacket) Read(b buffer.PacketBuffer) {
	packet.Index = b.ReadByte(b.Bytes(), b.Index())

}

func (packet *SelectMinigamesPacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.Index, b.Index())
}
