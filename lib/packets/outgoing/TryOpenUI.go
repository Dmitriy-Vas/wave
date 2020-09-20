package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *TryOpenUIPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *TryOpenUIPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *TryOpenUIPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *TryOpenUIPacket) SetSend(value bool) {
	packet.Send = value
}

type TryOpenUIPacket struct {
	ID      int64
	Send    bool
	UiIndex int32
}

func (packet *TryOpenUIPacket) Read(b buffer.PacketBuffer) {
	packet.UiIndex = b.ReadInt(b.Bytes(), b.Index())

}

func (packet *TryOpenUIPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.UiIndex, b.Index())
}
