package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type MapWornEquipmentPacket struct {
	*wave.DefaultPacket
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
