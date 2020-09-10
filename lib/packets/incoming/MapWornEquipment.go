package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *MapWornEquipmentPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *MapWornEquipmentPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *MapWornEquipmentPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *MapWornEquipmentPacket) SetSend(value bool) {
	d.Send = value
}

type MapWornEquipmentPacket struct {
	ID        int64
	Send      bool
	Variable1 int32
	Variable2 int32
}

func (packet *MapWornEquipmentPacket) Read(b buffer.PacketBuffer) {
	packet.Variable1 = b.ReadInt(b.Bytes(), b.Index())
	// TODO player equip data
	packet.Variable2 = b.ReadInt(b.Bytes(), b.Index())

}

func (packet *MapWornEquipmentPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Variable1, b.Index())
	// TODO player equip data
	b.WriteInt(b.Bytes(), packet.Variable2, b.Index())
}
