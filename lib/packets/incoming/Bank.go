package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib"
)

// GetID returns packet ID.
func (packet *BankPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *BankPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *BankPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *BankPacket) SetSend(value bool) {
	packet.Send = value
}

type BankPacket struct {
	ID        int64
	Send      bool
	Variable1 int32
	Slot      byte
	Items     []lib.InvItemRec
}

func (packet *BankPacket) Read(b buffer.PacketBuffer) {
	packet.Variable1 = b.ReadInt(b.Bytes(), b.Index())
	packet.Slot = b.ReadByte(b.Bytes(), b.Index())
	packet.Items = make([]lib.InvItemRec, 48+packet.Slot)
	if packet.Slot != 0 {
		packet.Items[0] = lib.ReadPlayerBankData(b)
	} else {
		for i := range packet.Items {
			packet.Items[i] = lib.ReadPlayerBankData(b)
		}
	}
}

func (packet *BankPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Variable1, b.Index())
	b.WriteByte(b.Bytes(), packet.Slot, b.Index())
}
