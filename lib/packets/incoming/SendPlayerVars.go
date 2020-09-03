package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type SendPlayerVarsPacket struct {
	*wave.DefaultPacket
	VarNum    int32
	Variables []bool
}

func (packet *SendPlayerVarsPacket) Read(b buffer.PacketBuffer) {
	packet.VarNum = b.ReadInt(b.Bytes(), b.Index())
	if packet.VarNum != 0 {
		packet.Variables = make([]bool, 1)
		packet.Variables[0] = b.ReadBool(b.Bytes(), b.Index())
	} else {
		packet.Variables = make([]bool, 100)
		for i, _ := range packet.Variables {
			packet.Variables[i] = b.ReadBool(b.Bytes(), b.Index())
		}
	}
}

func (packet *SendPlayerVarsPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.VarNum, b.Index())
	if packet.VarNum != 0 {
		b.WriteBool(b.Bytes(), packet.Variables[0], b.Index())
	} else {
		for _, v := range packet.Variables {
			b.WriteBool(b.Bytes(), v, b.Index())
		}
	}
}
