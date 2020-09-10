package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *PlayerMessagePacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *PlayerMessagePacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *PlayerMessagePacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *PlayerMessagePacket) SetSend(value bool) {
	d.Send = value
}

type PlayerMessagePacket struct {
	ID        int64
	Send      bool
	Variable0 bool
	TextNum   int32
	Messages  []string
	Message   string
	Color     string
	Variable1 bool
	ItemNum   int32
}

func (packet *PlayerMessagePacket) Read(b buffer.PacketBuffer) {
	if packet.Variable0 = b.ReadBool(b.Bytes(), b.Index()); !packet.Variable0 {
		packet.TextNum = b.ReadInt(b.Bytes(), b.Index())
		if amount := b.ReadInt(b.Bytes(), b.Index()); amount > -1 {
			packet.Messages = make([]string, amount)
			for i := range packet.Messages {
				packet.Messages[i] = b.ReadString(b.Bytes(), b.Index(), 0)
			}
		}
	} else {
		packet.Message = b.ReadString(b.Bytes(), b.Index(), 0)
	}
	packet.Color = b.ReadString(b.Bytes(), b.Index(), 0)
	if packet.Variable1 = b.ReadBool(b.Bytes(), b.Index()); packet.Variable1 {
		packet.ItemNum = b.ReadInt(b.Bytes(), b.Index())
		// TODO
		// if (modGameEditors.GetItemInt(invItemRec.Num, modEnumerations.MyItemInt.Enhancement) > 0)
		//{
		//	invItemRec.Slot = byteBuffer.ReadByte(true);
		//	invItemRec.Stat = new byte[6];
		//	int num3 = 0;
		//	do
		//	{
		//		invItemRec.Stat[num3] = byteBuffer.ReadByte(true);
		//		num3++;
		//	}
		//	while (num3 <= 5);
		//}
	}
	b.ReadByte(b.Bytes(), b.Index())
}

func (packet *PlayerMessagePacket) Write(b buffer.PacketBuffer) {

}
