package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib/objects"
)

type TakeCapturePacket struct {
	*wave.DefaultPacket
	Pos                  objects.Vector2
	NpcNum               int32
	InvestigationInvSlot byte
}

func (packet *TakeCapturePacket) Read(b buffer.PacketBuffer) {
	packet.Pos = objects.Vector2{X: b.ReadInt(b.Bytes(), b.Index()), Y: b.ReadInt(b.Bytes(), b.Index())}
	packet.NpcNum = b.ReadInt(b.Bytes(), b.Index())
	packet.InvestigationInvSlot = b.ReadByte(b.Bytes(), b.Index())

}

func (packet *TakeCapturePacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Pos.X, b.Index())
	b.WriteInt(b.Bytes(), packet.Pos.Y, b.Index())
	b.WriteInt(b.Bytes(), packet.NpcNum, b.Index())
	b.WriteByte(b.Bytes(), packet.InvestigationInvSlot, b.Index())
}
