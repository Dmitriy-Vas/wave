package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *ModifierPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *ModifierPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *ModifierPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *ModifierPacket) SetSend(value bool) {
	packet.Send = value
}

type ModifierPacket struct {
	ID         int64
	Send       bool
	Variable0  int64
	Variable1  float64
	Variable2  int32
	Variable3  bool
	Variable4  string
	Variable5  int32
	Variable6  int32
	Variable7  string
	Variable8  string
	Variable9  int32
	Variable10 float64
}

func (packet *ModifierPacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable1 = b.ReadDouble(b.Bytes(), b.Index())
	packet.Variable2 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable3 = b.ReadBool(b.Bytes(), b.Index())
	packet.Variable4 = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Variable5 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable6 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable7 = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Variable8 = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Variable9 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable10 = b.ReadDouble(b.Bytes(), b.Index())
}

func (packet *ModifierPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
	b.WriteDouble(b.Bytes(), packet.Variable1, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable2, b.Index())
	b.WriteBool(b.Bytes(), packet.Variable3, b.Index())
	b.WriteString(b.Bytes(), packet.Variable4, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable5, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable6, b.Index())
	b.WriteString(b.Bytes(), packet.Variable7, b.Index())
	b.WriteString(b.Bytes(), packet.Variable8, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable9, b.Index())
	b.WriteDouble(b.Bytes(), packet.Variable10, b.Index())
}
