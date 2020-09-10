package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *PlayerConnectedPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *PlayerConnectedPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *PlayerConnectedPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *PlayerConnectedPacket) SetSend(value bool) {
	d.Send = value
}

type PlayerConnectedPacket struct {
	ID              int64
	Send            bool
	Amount          int32
	SeveralMessages bool
	Messages        []string
}

func (packet *PlayerConnectedPacket) Read(b buffer.PacketBuffer) {
	packet.Amount = b.ReadInt(b.Bytes(), b.Index())
	packet.SeveralMessages = b.ReadBool(b.Bytes(), b.Index())
	if packet.SeveralMessages {
		packet.Messages = make([]string, packet.Amount)
		for i, _ := range packet.Messages {
			packet.Messages[i] = b.ReadString(b.Bytes(), b.Index(), 0)
		}
	}
}

func (packet *PlayerConnectedPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Amount, b.Index())
	b.WriteBool(b.Bytes(), packet.SeveralMessages, b.Index())
	if packet.SeveralMessages {
		for _, text := range packet.Messages {
			b.WriteString(b.Bytes(), text, b.Index())
		}
	}
}
