package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *UnequipPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *UnequipPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *UnequipPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *UnequipPacket) SetSend(value bool) {
	d.Send = value
}

type UnequipPacket struct {
	ID        int64
	Send      bool
	Cash      bool
	EquipSlot byte
}

func (packet *UnequipPacket) Read(b buffer.PacketBuffer) {
	packet.Cash = b.ReadBool(b.Bytes(), b.Index())
	packet.EquipSlot = b.ReadByte(b.Bytes(), b.Index())

}

func (packet *UnequipPacket) Write(b buffer.PacketBuffer) {
	b.WriteBool(b.Bytes(), packet.Cash, b.Index())
	b.WriteByte(b.Bytes(), packet.EquipSlot, b.Index())
}
