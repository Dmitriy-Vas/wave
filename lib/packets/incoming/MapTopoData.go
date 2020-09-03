package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib"
)

type MapTopoDataPacket struct {
	*wave.DefaultPacket
	Once    bool
	TopoNum []int32
	Topo    []lib.MapTopoRec
}

func (packet *MapTopoDataPacket) Read(b buffer.PacketBuffer) {
	amount := 500 // TODO move int to const
	if packet.Once = b.ReadBool(b.Bytes(), b.Index()); packet.Once {
		amount -= 499
	}
	packet.Topo = make([]lib.MapTopoRec, amount)
	packet.TopoNum = make([]int32, amount)
	for i := range packet.Topo {
		packet.TopoNum[i] = b.ReadInt(b.Bytes(), b.Index())
		topo := lib.MapTopoRec{
			Name: b.ReadString(b.Bytes(), b.Index(), 0),
		}
		if topo.Name != "" {
			topo.Item = make([]int64, 10)    // TODO move int to const
			topo.ItemVal = make([]int64, 10) // TODO move int to const
			topo.Luck = make([]int64, 10)    // TODO move int to const
			topo.Depth = int64(b.ReadInt(b.Bytes(), b.Index()))
			for x := range topo.Item {
				topo.Item[x] = int64(b.ReadInt(b.Bytes(), b.Index()))
				topo.ItemVal[x] = int64(b.ReadInt(b.Bytes(), b.Index()))
				topo.Luck[x] = int64(b.ReadInt(b.Bytes(), b.Index()))
			}
		}
		packet.Topo[i] = topo
	}
}

func (packet *MapTopoDataPacket) Write(b buffer.PacketBuffer) {
	b.WriteBool(b.Bytes(), packet.Once, b.Index())
	for i, topo := range packet.Topo {
		b.WriteInt(b.Bytes(), packet.TopoNum[i], b.Index())
		b.WriteString(b.Bytes(), topo.Name, b.Index())
		if topo.Name != "" {
			b.WriteInt(b.Bytes(), int32(topo.Depth), b.Index())
			for x := range topo.Item {
				b.WriteInt(b.Bytes(), int32(topo.Item[x]), b.Index())
				b.WriteInt(b.Bytes(), int32(topo.ItemVal[x]), b.Index())
				b.WriteInt(b.Bytes(), int32(topo.Luck[x]), b.Index())
			}
		}
	}
}
