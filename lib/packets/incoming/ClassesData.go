package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *ClassesDataPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *ClassesDataPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *ClassesDataPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *ClassesDataPacket) SetSend(value bool) {
	d.Send = value
}

type ClassesDataPacket struct {
	ID         int64
	Send       bool
	Variable0  int64
	Variable1  byte
	Variable2  string
	Variable3  string
	Variable4  string
	Variable5  int64
	Variable6  int32
	Variable7  string
	Variable8  string
	Variable9  bool
	Variable10 int32
	Variable11 string
	Variable12 bool
	Variable13 int32
	Variable14 int32
	Variable15 int32
	Variable16 string
	Variable17 bool
	Variable18 int32
	Variable19 int32
	Variable20 int32
}

func (packet *ClassesDataPacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable1 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable2 = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Variable3 = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Variable4 = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Variable5 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable6 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable7 = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Variable8 = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Variable9 = b.ReadBool(b.Bytes(), b.Index())
	packet.Variable10 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable11 = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Variable12 = b.ReadBool(b.Bytes(), b.Index())
	packet.Variable13 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable14 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable15 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable16 = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Variable17 = b.ReadBool(b.Bytes(), b.Index())
	packet.Variable18 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable19 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable20 = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *ClassesDataPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable1, b.Index())
	b.WriteString(b.Bytes(), packet.Variable2, b.Index())
	b.WriteString(b.Bytes(), packet.Variable3, b.Index())
	b.WriteString(b.Bytes(), packet.Variable4, b.Index())
	b.WriteLong(b.Bytes(), packet.Variable5, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable6, b.Index())
	b.WriteString(b.Bytes(), packet.Variable7, b.Index())
	b.WriteString(b.Bytes(), packet.Variable8, b.Index())
	b.WriteBool(b.Bytes(), packet.Variable9, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable10, b.Index())
	b.WriteString(b.Bytes(), packet.Variable11, b.Index())
	b.WriteBool(b.Bytes(), packet.Variable12, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable13, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable14, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable15, b.Index())
	b.WriteString(b.Bytes(), packet.Variable16, b.Index())
	b.WriteBool(b.Bytes(), packet.Variable17, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable18, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable19, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable20, b.Index())
}
