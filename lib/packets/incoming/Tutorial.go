package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *TutorialPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *TutorialPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *TutorialPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *TutorialPacket) SetSend(value bool) {
	packet.Send = value
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
