package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *HotbarUsePacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *HotbarUsePacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *HotbarUsePacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *HotbarUsePacket) SetSend(value bool) {
	packet.Send = value
}

type HotbarUsePacket struct {
	ID   int64
	Send bool
	Slot byte
}

func (packet *HotbarUsePacket) Read(b buffer.PacketBuffer) {
	packet.Slot = b.ReadByte(b.Bytes(), b.Index())

}

func (packet *HotbarUsePacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.Slot, b.Index())
}
