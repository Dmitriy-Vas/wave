package ROTMGWrapper

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

type Writer struct {
	*buffer.DefaultWriter
}

func InitWriter(writerInterface buffer.PacketWriter) {
	writer := writerInterface.(*Writer)
	writer.DefaultWriter = new(buffer.DefaultWriter)
	buffer.InitWriter(writer.DefaultWriter)
}
