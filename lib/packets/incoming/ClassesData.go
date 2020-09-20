package incoming

import (
	"fmt"
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib"
	"strconv"
	"strings"
)

// GetID returns packet ID.
func (packet *ClassesDataPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *ClassesDataPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *ClassesDataPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *ClassesDataPacket) SetSend(value bool) {
	packet.Send = value
}

type ClassesDataPacket struct {
	ID         int64
	Send       bool
	MaxClasses byte
	Classes    []lib.ClassRec
}

func (packet *ClassesDataPacket) Read(b buffer.PacketBuffer) {
	packet.MaxClasses = b.ReadByte(b.Bytes(), b.Index())
	packet.Classes = make([]lib.ClassRec, packet.MaxClasses)
	for i := range packet.Classes {
		packet.Classes[i].Stat = make([]byte, 6)   // TODO int to const
		packet.Classes[i].Vital = make([]int64, 3) // TODO int to const
		packet.Classes[i].Name = b.ReadString(b.Bytes(), b.Index(), 0)
		MFace := strings.Split(b.ReadString(b.Bytes(), b.Index(), 0), ",")
		FFace := strings.Split(b.ReadString(b.Bytes(), b.Index(), 0), ",")
		for x := range packet.Classes[i].Vital {
			packet.Classes[i].Vital[x] = b.ReadLong(b.Bytes(), b.Index())
		}
		hairTintAmount := b.ReadInt(b.Bytes(), b.Index())
		packet.Classes[i].HairTint = make([]lib.ClassHairTintRec, hairTintAmount)
		for x := range packet.Classes[i].HairTint {
			packet.Classes[i].HairTint[x].Name = b.ReadString(b.Bytes(), b.Index(), 0)
			packet.Classes[i].HairTint[x].Color = b.ReadString(b.Bytes(), b.Index(), 0)
			packet.Classes[i].HairTint[x].Premium = b.ReadBool(b.Bytes(), b.Index())
		}
		classSpriteAmount := b.ReadInt(b.Bytes(), b.Index())
		packet.Classes[i].Sprite = make([]lib.ClassSpriteRec, classSpriteAmount)
		for x := range packet.Classes[i].Sprite {
			packet.Classes[i].Sprite[x].Name = b.ReadString(b.Bytes(), b.Index(), 0)
			packet.Classes[i].Sprite[x].Premium = b.ReadBool(b.Bytes(), b.Index())
			packet.Classes[i].Sprite[x].Male = b.ReadInt(b.Bytes(), b.Index())
			packet.Classes[i].Sprite[x].Female = b.ReadInt(b.Bytes(), b.Index())
			if x < len(MFace) {
				if value, err := strconv.ParseInt(MFace[x], 10, 64); err == nil {
					packet.Classes[i].Sprite[x].MFace = int32(value)
				}
			}
			if x < len(FFace) {
				if value, err := strconv.ParseInt(FFace[x], 10, 64); err == nil {
					packet.Classes[i].Sprite[x].FFace = int32(value)
				}
			}
		}
		classHairAmount := b.ReadInt(b.Bytes(), b.Index())
		packet.Classes[i].Hair = make([]lib.ClassSpriteRec, classHairAmount)
		for x := range packet.Classes[i].Hair {
			packet.Classes[i].Hair[x].Name = b.ReadString(b.Bytes(), b.Index(), 0)
			packet.Classes[i].Hair[x].Premium = b.ReadBool(b.Bytes(), b.Index())
			packet.Classes[i].Hair[x].Male = b.ReadInt(b.Bytes(), b.Index())
			packet.Classes[i].Hair[x].Female = b.ReadInt(b.Bytes(), b.Index())
		}
		for x := range packet.Classes[i].Stat {
			packet.Classes[i].Stat[x] = byte(b.ReadInt(b.Bytes(), b.Index()))
		}
	}
}

func (packet *ClassesDataPacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.MaxClasses, b.Index())
	for _, class := range packet.Classes {
		b.WriteString(b.Bytes(), class.Name, b.Index())
		MFace := make([]string, 0)
		FFace := make([]string, 0)
		for _, sprite := range class.Sprite {
			if sprite.MFace != 0 {
				MFace = append(MFace, fmt.Sprint(sprite.MFace))
			}
			if sprite.FFace != 0 {
				FFace = append(FFace, fmt.Sprint(sprite.FFace))
			}
		}
		b.WriteString(b.Bytes(), strings.Join(MFace, ","), b.Index())
		b.WriteString(b.Bytes(), strings.Join(FFace, ","), b.Index())
		for _, vital := range class.Vital {
			b.WriteLong(b.Bytes(), vital, b.Index())
		}
		b.WriteInt(b.Bytes(), int32(len(class.HairTint)), b.Index())
		for _, tint := range class.HairTint {
			b.WriteString(b.Bytes(), tint.Name, b.Index())
			b.WriteString(b.Bytes(), tint.Color, b.Index())
			b.WriteBool(b.Bytes(), tint.Premium, b.Index())
		}
		b.WriteInt(b.Bytes(), int32(len(class.Sprite)), b.Index())
		for _, sprite := range class.Sprite {
			b.WriteString(b.Bytes(), sprite.Name, b.Index())
			b.WriteBool(b.Bytes(), sprite.Premium, b.Index())
			b.WriteInt(b.Bytes(), sprite.Male, b.Index())
			b.WriteInt(b.Bytes(), sprite.Female, b.Index())
		}
		b.WriteInt(b.Bytes(), int32(len(class.Hair)), b.Index())
		for _, hair := range class.Hair {
			b.WriteString(b.Bytes(), hair.Name, b.Index())
			b.WriteBool(b.Bytes(), hair.Premium, b.Index())
			b.WriteInt(b.Bytes(), hair.Male, b.Index())
			b.WriteInt(b.Bytes(), hair.Female, b.Index())
		}
		for _, stat := range class.Stat {
			b.WriteInt(b.Bytes(), int32(stat), b.Index())
		}
	}
}
