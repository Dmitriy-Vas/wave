package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type ReceiveHourPacket struct {
	*wave.DefaultPacket
	Hour        int32
	ServerDay   byte
	ServerMonth byte
}

func (packet *ReceiveHourPacket) Read(b buffer.PacketBuffer) {
	packet.Hour = b.ReadInt(b.Bytes(), b.Index())
	packet.ServerDay = b.ReadByte(b.Bytes(), b.Index())
	packet.ServerMonth = b.ReadByte(b.Bytes(), b.Index())
}

func (packet *ReceiveHourPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Hour, b.Index())
	b.WriteByte(b.Bytes(), packet.ServerDay, b.Index())
	b.WriteByte(b.Bytes(), packet.ServerMonth, b.Index())
}
