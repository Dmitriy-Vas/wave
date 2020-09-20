package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *SendConditionVarPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *SendConditionVarPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *SendConditionVarPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *SendConditionVarPacket) SetSend(value bool) {
	packet.Send = value
}

type SendConditionVarPacket struct {
	ID            int64
	Send          bool
	ConditionVars []string
}

func (packet *SendConditionVarPacket) Read(b buffer.PacketBuffer) {
	packet.ConditionVars = make([]string, 100)
	for i := range packet.ConditionVars {
		packet.ConditionVars[i] = b.ReadString(b.Bytes(), b.Index(), 0)
	}
}

func (packet *SendConditionVarPacket) Write(b buffer.PacketBuffer) {
	for _, condition := range packet.ConditionVars {
		b.WriteString(b.Bytes(), condition, b.Index())
	}
}
