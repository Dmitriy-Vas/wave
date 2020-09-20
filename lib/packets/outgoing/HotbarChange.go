package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *HotbarChangePacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *HotbarChangePacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *HotbarChangePacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *HotbarChangePacket) SetSend(value bool) {
	packet.Send = value
}

type HotbarChangePacket struct {
	ID        int64
	Send      bool
	Type      byte
	Slot      int32
	HotbarNum int32
}

func (packet *HotbarChangePacket) Read(b buffer.PacketBuffer) {
	packet.Type = b.ReadByte(b.Bytes(), b.Index())
	packet.Slot = b.ReadInt(b.Bytes(), b.Index())
	packet.HotbarNum = b.ReadInt(b.Bytes(), b.Index())

}

func (packet *HotbarChangePacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.Type, b.Index())
	b.WriteInt(b.Bytes(), packet.Slot, b.Index())
	b.WriteInt(b.Bytes(), packet.HotbarNum, b.Index())
}
