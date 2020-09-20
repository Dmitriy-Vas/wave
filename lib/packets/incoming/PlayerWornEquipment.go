package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *PlayerWornEquipmentPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *PlayerWornEquipmentPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *PlayerWornEquipmentPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *PlayerWornEquipmentPacket) SetSend(value bool) {
	packet.Send = value
}

type PlayerWornEquipmentPacket struct {
	ID   int64
	Send bool
}

func (packet *PlayerWornEquipmentPacket) Read(_ buffer.PacketBuffer) {

}

func (packet *PlayerWornEquipmentPacket) Write(_ buffer.PacketBuffer) {

}
