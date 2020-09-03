package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type SendConditionVarPacket struct {
	*wave.DefaultPacket
	ConditionVars []string
}

func (packet *SendConditionVarPacket) Read(b buffer.PacketBuffer) {
	packet.ConditionVars = make([]string, 100)
	for i, _ := range packet.ConditionVars {
		packet.ConditionVars[i] = b.ReadString(b.Bytes(), b.Index(), 0)
	}
}

func (packet *SendConditionVarPacket) Write(b buffer.PacketBuffer) {
	for _, condition := range packet.ConditionVars {
		b.WriteString(b.Bytes(), condition, b.Index())
	}
}
