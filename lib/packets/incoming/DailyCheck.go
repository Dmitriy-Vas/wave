package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib/wrapper"
	"time"
)

// GetID returns packet ID.
func (d *DailyCheckPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *DailyCheckPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *DailyCheckPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *DailyCheckPacket) SetSend(value bool) {
	d.Send = value
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
