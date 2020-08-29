package DSIncoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/packets"
)

type RequestEditorPacket struct {
	*packets.DefaultPacket
	Variable0 int64
	Variable1 int32
}

func init() {
	packets.ServerPackets = append(packets.ServerPackets, &RequestEditorPacket{
		DefaultPacket: new(packets.DefaultPacket),
	})
}

func (packet *RequestEditorPacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable1 = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *RequestEditorPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable1, b.Index())
}
