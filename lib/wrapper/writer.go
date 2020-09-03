package wrapper

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"time"
)

type Writer struct {
	*buffer.DefaultWriter
}

func InitWriter(writerInterface buffer.PacketWriter) {
	writer := writerInterface.(*Writer)
	writer.DefaultWriter = new(buffer.DefaultWriter)
	buffer.InitWriter(writer.DefaultWriter)
}

func WriteDate(buffer buffer.PacketBuffer, t time.Time) {
	str := t.Format("02/01/2006")
	buffer.WriteString(buffer.Bytes(), str, buffer.Index())
}
