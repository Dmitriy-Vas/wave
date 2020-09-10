package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *TryOpenUIPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *TryOpenUIPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *TryOpenUIPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *TryOpenUIPacket) SetSend(value bool) {
	d.Send = value
}

type TryOpenUIPacket struct {
	ID      int64
	Send    bool
	UiIndex int32
}

func (packet *TryOpenUIPacket) Read(b buffer.PacketBuffer) {
	packet.UiIndex = b.ReadInt(b.Bytes(), b.Index())

}

func (packet *TryOpenUIPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.UiIndex, b.Index())
}
