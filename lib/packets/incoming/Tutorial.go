package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *TutorialPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *TutorialPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *TutorialPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *TutorialPacket) SetSend(value bool) {
	d.Send = value
}

type TutorialPacket struct {
	ID          int64
	Send        bool
	TutorialNum int32
}

func (packet *TutorialPacket) Read(b buffer.PacketBuffer) {
	packet.TutorialNum = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *TutorialPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.TutorialNum, b.Index())
}
