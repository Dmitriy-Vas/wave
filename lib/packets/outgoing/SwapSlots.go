package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *SwapSlotsPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *SwapSlotsPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *SwapSlotsPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *SwapSlotsPacket) SetSend(value bool) {
	d.Send = value
}

type SwapSlotsPacket struct {
	ID      int64
	Send    bool
	OldSlot int32
	NewSlot int32
	Type    byte
}

func (packet *SwapSlotsPacket) Read(b buffer.PacketBuffer) {
	packet.OldSlot = b.ReadInt(b.Bytes(), b.Index())
	packet.NewSlot = b.ReadInt(b.Bytes(), b.Index())
	packet.Type = b.ReadByte(b.Bytes(), b.Index())

}

func (packet *SwapSlotsPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.OldSlot, b.Index())
	b.WriteInt(b.Bytes(), packet.NewSlot, b.Index())
	b.WriteByte(b.Bytes(), packet.Type, b.Index())
}
