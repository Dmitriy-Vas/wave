package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *FlashPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *FlashPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *FlashPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *FlashPacket) SetSend(value bool) {
	packet.Send = value
}

type FlashPacket struct {
	ID         int64
	Send       bool
	Variable1  byte
	Variable2  int32
	Variable3  byte
	Expression bool
	Num        int32
}

func (packet *FlashPacket) Read(b buffer.PacketBuffer) {
	packet.Variable1 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable2 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable3 = b.ReadByte(b.Bytes(), b.Index())
	packet.Expression = b.ReadBool(b.Bytes(), b.Index())
	packet.Num = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *FlashPacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.Variable1, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable2, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable3, b.Index())
	b.WriteBool(b.Bytes(), packet.Expression, b.Index())
	b.WriteInt(b.Bytes(), packet.Num, b.Index())
}
