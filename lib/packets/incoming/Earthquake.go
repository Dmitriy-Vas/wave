package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *EarthquakePacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *EarthquakePacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *EarthquakePacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *EarthquakePacket) SetSend(value bool) {
	packet.Send = value
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
