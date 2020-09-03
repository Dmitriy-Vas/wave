package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type ServerVarsPacket struct {
	*wave.DefaultPacket
	GlobalVar int32
}

func (packet *ServerVarsPacket) Read(b buffer.PacketBuffer) {
	packet.GlobalVar = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *ServerVarsPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.GlobalVar, b.Index())
}
