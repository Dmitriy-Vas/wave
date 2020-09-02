package wrapper

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

type Reader struct {
	*buffer.DefaultReader
}

func InitReader(readerInterface buffer.PacketReader) {
	reader := readerInterface.(*Reader)
	reader.DefaultReader = new(buffer.DefaultReader)
	buffer.InitReader(reader.DefaultReader)
}
