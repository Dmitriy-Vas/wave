package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *UpdateConditionPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *UpdateConditionPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *UpdateConditionPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *UpdateConditionPacket) SetSend(value bool) {
	d.Send = value
}

type UpdateConditionPacket struct {
	ID        int64
	Send      bool
	Variable0 int32
}

func (packet *UpdateConditionPacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadInt(b.Bytes(), b.Index())
	if v := packet.Variable0; v != 0 || v <= 400 {
		// TODO condition data
	}
}

func (packet *UpdateConditionPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Variable0, b.Index())
	if v := packet.Variable0; v != 0 || v <= 400 {
		// TODO condition data
	}
}
