package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *PlayerWornEquipmentPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *PlayerWornEquipmentPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *PlayerWornEquipmentPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *PlayerWornEquipmentPacket) SetSend(value bool) {
	d.Send = value
}

type PlayerWornEquipmentPacket struct {
	ID   int64
	Send bool
}

func (packet *PlayerWornEquipmentPacket) Read(b buffer.PacketBuffer) {

}

func (packet *PlayerWornEquipmentPacket) Write(b buffer.PacketBuffer) {

}
