package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib/wrapper"
	"time"
)

// GetID returns packet ID.
func (packet *DailyCheckPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *DailyCheckPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *DailyCheckPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *DailyCheckPacket) SetSend(value bool) {
	packet.Send = value
}

type DailyCheckPacket struct {
	ID              int64
	Send            bool
	CheckPlayerDate time.Time
	CheckCount      byte
	CheckDate       time.Time
}

func (packet *DailyCheckPacket) Read(b buffer.PacketBuffer) {
	packet.CheckPlayerDate = wrapper.ReadDate(b)
	packet.CheckCount = b.ReadByte(b.Bytes(), b.Index())
	packet.CheckDate = wrapper.ReadDate(b)

}

func (packet *DailyCheckPacket) Write(b buffer.PacketBuffer) {
	wrapper.WriteDate(b, packet.CheckPlayerDate)
	b.WriteByte(b.Bytes(), packet.CheckCount, b.Index())
	wrapper.WriteDate(b, packet.CheckDate)
}
