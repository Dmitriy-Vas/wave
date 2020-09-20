package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *UseMegaphonePacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *UseMegaphonePacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *UseMegaphonePacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *UseMegaphonePacket) SetSend(value bool) {
	packet.Send = value
}

type UseMegaphonePacket struct {
	ID      int64
	Send    bool
	Type    byte
	Message string
	InvItem byte
}

func (packet *UseMegaphonePacket) Read(b buffer.PacketBuffer) {
	packet.Type = b.ReadByte(b.Bytes(), b.Index())
	packet.Message = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.InvItem = b.ReadByte(b.Bytes(), b.Index())

}

func (packet *UseMegaphonePacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.Type, b.Index())
	b.WriteString(b.Bytes(), packet.Message, b.Index())
	b.WriteByte(b.Bytes(), packet.InvItem, b.Index())
}
