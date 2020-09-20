package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *TrainStatPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *TrainStatPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *TrainStatPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *TrainStatPacket) SetSend(value bool) {
	packet.Send = value
}

type TrainStatPacket struct {
	ID                 int64
	Send               bool
	StatType           byte
	IsCharacter        bool
	Value              int32
	SelectedProfession byte
}

func (packet *TrainStatPacket) Read(b buffer.PacketBuffer) {
	packet.StatType = b.ReadByte(b.Bytes(), b.Index())
	packet.IsCharacter = b.ReadBool(b.Bytes(), b.Index())
	packet.Value = b.ReadInt(b.Bytes(), b.Index())
	packet.SelectedProfession = b.ReadByte(b.Bytes(), b.Index())

}

func (packet *TrainStatPacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.StatType, b.Index())
	b.WriteBool(b.Bytes(), packet.IsCharacter, b.Index())
	b.WriteInt(b.Bytes(), packet.Value, b.Index())
	b.WriteByte(b.Bytes(), packet.SelectedProfession, b.Index())
}
