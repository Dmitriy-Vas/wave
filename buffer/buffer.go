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
	Clone() PacketBuffer
}

// Buffer capacity sets once and can't grow up
// To prevent copying data on each buffer resizing
type DefaultBuffer struct {
	PacketReader
	PacketWriter

	init_len uint64
	max_len  uint64
	len      uint64
	index    uint64
	buf      []byte
}

func (db *DefaultBuffer) Clone() PacketBuffer {
	return &DefaultBuffer{
		PacketReader: db.PacketReader,
		PacketWriter: db.PacketWriter,
		init_len:     db.init_len,
		max_len:      db.max_len,
		buf:          make([]byte, db.max_len),
	}
}

func (db *DefaultBuffer) SetInitLength(length uint64) {
	db.init_len = length
}

func (db *DefaultBuffer) SetMaxLength(size uint64) {
	db.max_len = size
	db.buf = make([]byte, size)
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
	if size := db.index + amount; db.len <= size {
		db.Resize(size)
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
	return db.buf
}
