package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *HotbarChangePacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *HotbarChangePacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *HotbarChangePacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *HotbarChangePacket) SetSend(value bool) {
	d.Send = value
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
