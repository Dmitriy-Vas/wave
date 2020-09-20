package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *ProfessionEndPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *ProfessionEndPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *ProfessionEndPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *ProfessionEndPacket) SetSend(value bool) {
	packet.Send = value
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
