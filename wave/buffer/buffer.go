package buffer

type PacketBuffer interface {
	PacketReader
	PacketWriter
	SetReader(reader PacketReader)
	SetWriter(writer PacketWriter)

	Resize(size uint64)
	Next(amount uint64) uint64
	Back(amount uint64) uint64
	Reset()
	Len() uint64
	Index() uint64
	Bytes() []byte
}

type Vector2 struct {
	X int32
	Y int32
}

// Buffer capacity sets once and can't grow up
// To exclude copying data on each buffer resizing
type DefaultBuffer struct {
	PacketReader
	PacketWriter

	init_len uint64
	max_len  uint64
	len      uint64
	index    uint64
	buf      []byte
}

func InitBuffer(bufferInterface PacketBuffer) {
	buffer := bufferInterface.(*DefaultBuffer)
	buffer.buf = make([]byte, buffer.max_len)
}

func (db *DefaultBuffer) SetInitialLength(size uint64) {
	db.init_len = size
}

func (db *DefaultBuffer) SetMaxLength(size uint64) {
	db.max_len = size
}

func (db *DefaultBuffer) SetReader(reader PacketReader) {
	db.PacketReader = reader
}

func (db *DefaultBuffer) SetWriter(writer PacketWriter) {
	db.PacketWriter = writer
}

func (db *DefaultBuffer) Resize(size uint64) {
	if size > db.max_len {
		panic("New buffer size is too large")
	}
	db.len = size
}

func (db *DefaultBuffer) Next(amount uint64) uint64 {
	if size := db.index + amount; db.len < size {
		db.len = size
	}
	db.index += amount
	return db.index
}

func (db *DefaultBuffer) Back(amount uint64) uint64 {
	db.index -= amount
	return db.index
}

func (db *DefaultBuffer) Reset() {
	db.index = 0
	db.len = db.init_len
}

func (db *DefaultBuffer) Len() uint64 {
	return db.len
}

func (db *DefaultBuffer) Index() uint64 {
	return db.index
}

func (db *DefaultBuffer) Bytes() []byte {
	return db.buf[:db.len]
}
