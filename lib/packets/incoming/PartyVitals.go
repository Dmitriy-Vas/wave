package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *PartyVitalsPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *PartyVitalsPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *PartyVitalsPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *PartyVitalsPacket) SetSend(value bool) {
	packet.Send = value
}

type PartyVitalsPacket struct {
	ID         int64
	Send       bool
	Variable0  int64
	Variable1  bool
	Variable2  int32
	Variable3  int64
	Variable4  int32
	Variable5  int64
	Variable6  int64
	Variable7  int32
	Variable8  int32
	Variable9  int32
	Variable10 int32
	Variable11 string
	Variable12 int32
	Variable13 int32
}

func (packet *PartyVitalsPacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable1 = b.ReadBool(b.Bytes(), b.Index())
	packet.Variable2 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable3 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable4 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable5 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable6 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable7 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable8 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable9 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable10 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable11 = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Variable12 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable13 = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *PartyVitalsPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
	b.WriteBool(b.Bytes(), packet.Variable1, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable2, b.Index())
	b.WriteLong(b.Bytes(), packet.Variable3, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable4, b.Index())
	b.WriteLong(b.Bytes(), packet.Variable5, b.Index())
	b.WriteLong(b.Bytes(), packet.Variable6, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable7, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable8, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable9, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable10, b.Index())
	b.WriteString(b.Bytes(), packet.Variable11, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable12, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable13, b.Index())
}
