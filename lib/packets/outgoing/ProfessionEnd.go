package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *ProfessionEndPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *ProfessionEndPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *ProfessionEndPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *ProfessionEndPacket) SetSend(value bool) {
	d.Send = value
}

type ProfessionEndPacket struct {
	ID           int64
	Send         bool
	FishingCount int32
	FishingPos   int32
	FishingFail  byte
	FishingTime  int32
}

func (packet *ProfessionEndPacket) Read(b buffer.PacketBuffer) {
	packet.FishingCount = b.ReadInt(b.Bytes(), b.Index())
	packet.FishingPos = b.ReadInt(b.Bytes(), b.Index())
	packet.FishingFail = b.ReadByte(b.Bytes(), b.Index())
	packet.FishingTime = b.ReadInt(b.Bytes(), b.Index())

}

func (packet *ProfessionEndPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.FishingCount, b.Index())
	b.WriteInt(b.Bytes(), packet.FishingPos, b.Index())
	b.WriteByte(b.Bytes(), packet.FishingFail, b.Index())
	b.WriteInt(b.Bytes(), packet.FishingTime, b.Index())
}
