package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *ButtonPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *ButtonPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *ButtonPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *ButtonPacket) SetSend(value bool) {
	packet.Send = value
}

type ButtonPacket struct {
	ID       int64
	Send     bool
	UiIndex  int32
	UiButton byte
	Value    int32
}

func (packet *ButtonPacket) Read(b buffer.PacketBuffer) {
	packet.UiIndex = b.ReadInt(b.Bytes(), b.Index())
	packet.UiButton = b.ReadByte(b.Bytes(), b.Index())
	packet.Value = b.ReadInt(b.Bytes(), b.Index())

}

func (packet *ButtonPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.UiIndex, b.Index())
	b.WriteByte(b.Bytes(), packet.UiButton, b.Index())
	b.WriteInt(b.Bytes(), packet.Value, b.Index())
}
