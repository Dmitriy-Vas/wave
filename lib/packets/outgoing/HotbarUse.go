package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *HotbarUsePacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *HotbarUsePacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *HotbarUsePacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *HotbarUsePacket) SetSend(value bool) {
	d.Send = value
}

type HotbarUsePacket struct {
	ID   int64
	Send bool
	Slot byte
}

func (packet *HotbarUsePacket) Read(b buffer.PacketBuffer) {
	packet.Slot = b.ReadByte(b.Bytes(), b.Index())

}

func (packet *HotbarUsePacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.Slot, b.Index())
}
