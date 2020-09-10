package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *CashInvUpdatePacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *CashInvUpdatePacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *CashInvUpdatePacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *CashInvUpdatePacket) SetSend(value bool) {
	d.Send = value
}

type CashInvUpdatePacket struct {
	ID        int64
	Send      bool
	Slot      byte
	ItemNum   int32
	ItemValue int32
}

func (packet *CashInvUpdatePacket) Read(b buffer.PacketBuffer) {
	packet.Slot = b.ReadByte(b.Bytes(), b.Index())
	b.ReadInt(b.Bytes(), b.Index())
	b.ReadInt(b.Bytes(), b.Index())
}

func (packet *CashInvUpdatePacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.Slot, b.Index())
	b.WriteInt(b.Bytes(), packet.ItemNum, b.Index())
	b.WriteInt(b.Bytes(), packet.ItemValue, b.Index())
}
