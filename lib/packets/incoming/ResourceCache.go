package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib"
)

type ResourceCachePacket struct {
	*wave.DefaultPacket
	ResourceIndex int32
	MapResources  []lib.MapResourceRec
}

func (packet *ResourceCachePacket) Read(b buffer.PacketBuffer) {
	if packet.ResourceIndex = b.ReadInt(b.Bytes(), b.Index()); packet.ResourceIndex > 0 {
		packet.MapResources = make([]lib.MapResourceRec, packet.ResourceIndex)
		for i, _ := range packet.MapResources {
			packet.MapResources[i] = lib.MapResourceRec{
				ResourceState: b.ReadByte(b.Bytes(), b.Index()),
				X:             b.ReadInt(b.Bytes(), b.Index()),
				Y:             b.ReadInt(b.Bytes(), b.Index()),
				Shadow:        b.ReadByte(b.Bytes(), b.Index()),
			}
		}
	}
}

func (packet *ResourceCachePacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.ResourceIndex, b.Index())
	for _, res := range packet.MapResources {
		b.WriteByte(b.Bytes(), res.ResourceState, b.Index())
		b.WriteInt(b.Bytes(), res.X, b.Index())
		b.WriteInt(b.Bytes(), res.Y, b.Index())
		b.WriteByte(b.Bytes(), res.Shadow, b.Index())
	}
}
