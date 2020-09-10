package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *TrainStatPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *TrainStatPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *TrainStatPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *TrainStatPacket) SetSend(value bool) {
	d.Send = value
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
