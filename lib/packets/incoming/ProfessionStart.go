package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *ProfessionStartPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *ProfessionStartPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *ProfessionStartPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *ProfessionStartPacket) SetSend(value bool) {
	d.Send = value
}

type ProfessionStartPacket struct {
	ID           int64
	Send         bool
	Variable1    byte
	FishingDiff  int32
	FishingNum   int32
	FishingSpeed int32
}

func (packet *ProfessionStartPacket) Read(b buffer.PacketBuffer) {
	packet.Variable1 = b.ReadByte(b.Bytes(), b.Index())
	packet.FishingDiff = b.ReadInt(b.Bytes(), b.Index())
	packet.FishingNum = b.ReadInt(b.Bytes(), b.Index())
	packet.FishingSpeed = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *ProfessionStartPacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.Variable1, b.Index())
	b.WriteInt(b.Bytes(), packet.FishingDiff, b.Index())
	b.WriteInt(b.Bytes(), packet.FishingNum, b.Index())
	b.WriteInt(b.Bytes(), packet.FishingSpeed, b.Index())
}
