package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type ButtonPacket struct {
	*wave.DefaultPacket
	UiIndex  int32
	UiButton byte
	Value    int32
}

func (packet *ButtonPacket) Read(b buffer.PacketBuffer) {
	packet.UiIndex = b.ReadInt(b.Bytes(), b.Index())
	packet.UiButton = b.ReadByte(b.Bytes(), b.Index())
	packet.Value = b.ReadInt(b.Bytes(), b.Index())

}

func (packet *ButtonPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.UiIndex, b.Index())
	b.WriteByte(b.Bytes(), packet.UiButton, b.Index())
	b.WriteInt(b.Bytes(), packet.Value, b.Index())
}
