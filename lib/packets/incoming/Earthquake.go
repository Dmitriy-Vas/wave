package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *EarthquakePacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *EarthquakePacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *EarthquakePacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *EarthquakePacket) SetSend(value bool) {
	d.Send = value
}

type EarthquakePacket struct {
	ID            int64
	Send          bool
	ShowEartquake byte
}

func (packet *EarthquakePacket) Read(b buffer.PacketBuffer) {
	packet.ShowEartquake = b.ReadByte(b.Bytes(), b.Index())
}

func (packet *EarthquakePacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.ShowEartquake, b.Index())
}
