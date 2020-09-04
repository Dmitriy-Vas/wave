package wrapper

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"time"
)

type Reader struct {
	*buffer.DefaultReader
}

func InitReader(readerInterface buffer.PacketReader) {
	reader := readerInterface.(*Reader)
	reader.DefaultReader = new(buffer.DefaultReader)
	buffer.InitReader(reader.DefaultReader)
}

func ReadDate(buffer buffer.PacketBuffer) time.Time {
	str := buffer.ReadString(buffer.Bytes(), buffer.Index(), 0)
	if str == "" {
		return time.Time{}
	}
	t, _ := time.Parse("02/01/2006", str)
	return t
}
