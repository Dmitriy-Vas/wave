package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type TrainStatPacket struct {
	*wave.DefaultPacket
	StatType           byte
	IsCharacter        bool
	Value              int32
	SelectedProfession byte
}

func (packet *TrainStatPacket) Read(b buffer.PacketBuffer) {
	packet.StatType = b.ReadByte(b.Bytes(), b.Index())
	packet.IsCharacter = b.ReadBool(b.Bytes(), b.Index())
	packet.Value = b.ReadInt(b.Bytes(), b.Index())
	packet.SelectedProfession = b.ReadByte(b.Bytes(), b.Index())

}

func (packet *TrainStatPacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.StatType, b.Index())
	b.WriteBool(b.Bytes(), packet.IsCharacter, b.Index())
	b.WriteInt(b.Bytes(), packet.Value, b.Index())
	b.WriteByte(b.Bytes(), packet.SelectedProfession, b.Index())
}
