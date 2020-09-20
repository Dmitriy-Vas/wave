package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *SendPlayerVarsPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *SendPlayerVarsPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *SendPlayerVarsPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *SendPlayerVarsPacket) SetSend(value bool) {
	packet.Send = value
}

type SendPlayerVarsPacket struct {
	ID        int64
	Send      bool
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
		for i := range packet.Variables {
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
