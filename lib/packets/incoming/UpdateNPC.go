package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type UpdateNPCPacket struct {
	*wave.DefaultPacket
	Num int32
}

func (packet *UpdateNPCPacket) Read(b buffer.PacketBuffer) {
	packet.Num = b.ReadInt(b.Bytes(), b.Index())
	// TODO NPC Data
}

func (packet *UpdateNPCPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Num, b.Index())
	// TODO NPC Data
}
