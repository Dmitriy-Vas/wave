package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type UnequipPacket struct {
	*wave.DefaultPacket
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
