package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *AddCharPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *AddCharPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *AddCharPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *AddCharPacket) SetSend(value bool) {
	packet.Send = value
}

type AddCharPacket struct {
	ID       int64
	Send     bool
	Name     string
	Sex      byte
	ClassNum int32
	Sprite   int32
	Hair     int32
	HairTint int32
	Language string
}

func (packet *AddCharPacket) Read(b buffer.PacketBuffer) {
	packet.Name = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Sex = b.ReadByte(b.Bytes(), b.Index())
	packet.ClassNum = b.ReadInt(b.Bytes(), b.Index())
	packet.Sprite = b.ReadInt(b.Bytes(), b.Index())
	packet.Hair = b.ReadInt(b.Bytes(), b.Index())
	packet.HairTint = b.ReadInt(b.Bytes(), b.Index())
	packet.Language = b.ReadString(b.Bytes(), b.Index(), 0)

}

func (packet *AddCharPacket) Write(b buffer.PacketBuffer) {
	b.WriteString(b.Bytes(), packet.Name, b.Index())
	b.WriteByte(b.Bytes(), packet.Sex, b.Index())
	b.WriteInt(b.Bytes(), packet.ClassNum, b.Index())
	b.WriteInt(b.Bytes(), packet.Sprite, b.Index())
	b.WriteInt(b.Bytes(), packet.Hair, b.Index())
	b.WriteInt(b.Bytes(), packet.HairTint, b.Index())
	b.WriteString(b.Bytes(), packet.Language, b.Index())
}
