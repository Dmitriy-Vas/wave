package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib"
)

// GetID returns packet ID.
func (d *CharDataPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *CharDataPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *CharDataPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *CharDataPacket) SetSend(value bool) {
	d.Send = value
}

type CharDataPacket struct {
	ID    int64
	Send  bool
	Type  []int32
	Chars []lib.CharDataRec
}

func (packet *CharDataPacket) Read(b buffer.PacketBuffer) {
	packet.Chars = make([]lib.CharDataRec, 3) // TODO move to constants
	packet.Type = make([]int32, 3)
	for i := range packet.Chars {
		if packet.Type[i] = b.ReadInt(b.Bytes(), b.Index()); packet.Type[i] > 0 {
			char := lib.CharDataRec{
				Name:      b.ReadString(b.Bytes(), b.Index(), 0),
				Level:     b.ReadInt(b.Bytes(), b.Index()),
				Classes:   int32(b.ReadByte(b.Bytes(), b.Index())),
				Sex:       b.ReadByte(b.Bytes(), b.Index()),
				Sprite:    b.ReadInt(b.Bytes(), b.Index()),
				Hair:      b.ReadInt(b.Bytes(), b.Index()),
				HairTint:  b.ReadString(b.Bytes(), b.Index(), 0),
				Equip:     make([]int32, 16), // TODO move to constants
				CashEquip: make([]int32, 12), // TODO move to constants
				Invisible: make([]bool, 12),  // TODO move to constants
			}
			for x := range char.Equip {
				char.Equip[x] = b.ReadInt(b.Bytes(), b.Index())
			}
			for x := range char.CashEquip {
				char.CashEquip[x] = b.ReadInt(b.Bytes(), b.Index())
				char.Invisible[x] = b.ReadBool(b.Bytes(), b.Index())
			}
			char.NotHair = b.ReadBool(b.Bytes(), b.Index())
			packet.Chars[i] = char
		}
	}
}

func (packet *CharDataPacket) Write(b buffer.PacketBuffer) {
	for i, char := range packet.Chars {
		b.WriteInt(b.Bytes(), packet.Type[i], b.Index())
		if packet.Type[i] > 0 {
			b.WriteString(b.Bytes(), char.Name, b.Index())
			b.WriteInt(b.Bytes(), char.Level, b.Index())
			b.WriteByte(b.Bytes(), byte(char.Classes), b.Index())
			b.WriteByte(b.Bytes(), char.Sex, b.Index())
			b.WriteInt(b.Bytes(), char.Sprite, b.Index())
			b.WriteInt(b.Bytes(), char.Hair, b.Index())
			b.WriteString(b.Bytes(), char.HairTint, b.Index())
			for _, e := range char.Equip {
				b.WriteInt(b.Bytes(), e, b.Index())
			}
			for x := range char.CashEquip {
				b.WriteInt(b.Bytes(), char.CashEquip[x], b.Index())
				b.WriteBool(b.Bytes(), char.Invisible[x], b.Index())
			}
			b.WriteBool(b.Bytes(), char.NotHair, b.Index())
		}
	}
}
