package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type PlayerWornEquipmentPacket struct {
	*wave.DefaultPacket
}

func (packet *PlayerWornEquipmentPacket) Read(b buffer.PacketBuffer) {

}

func (packet *PlayerWornEquipmentPacket) Write(b buffer.PacketBuffer) {

}
