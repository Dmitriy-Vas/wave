package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *SendConditionVarPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *SendConditionVarPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *SendConditionVarPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *SendConditionVarPacket) SetSend(value bool) {
	d.Send = value
}

type SendConditionVarPacket struct {
	ID            int64
	Send          bool
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
