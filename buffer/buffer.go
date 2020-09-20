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

	initLen uint64
	maxLen  uint64
	len     uint64
	index   uint64
	buf     []byte
}

func (db *DefaultBuffer) WriteBool(data []byte, value bool, index uint64) {
	db.Next(1)
	db.PacketWriter.WriteBool(data, value, index)
}

func (db *DefaultBuffer) WriteByte(data []byte, value byte, index uint64) {
	db.Next(1)
	db.PacketWriter.WriteByte(data, value, index)
}

func (db *DefaultBuffer) WriteFloat(data []byte, value float32, index uint64) {
	db.Next(4)
	db.PacketWriter.WriteFloat(data, value, index)
}

func (db *DefaultBuffer) WriteDouble(data []byte, value float64, index uint64) {
	db.Next(8)
	db.PacketWriter.WriteDouble(data, value, index)
}

func (db *DefaultBuffer) WriteInt(data []byte, value int32, index uint64) {
	db.Next(4)
	db.PacketWriter.WriteInt(data, value, index)
}

func (db *DefaultBuffer) WriteUint(data []byte, value uint32, index uint64) {
	db.Next(4)
	db.PacketWriter.WriteUint(data, value, index)
}

func (db *DefaultBuffer) WriteLong(data []byte, value int64, index uint64) {
	db.Next(8)
	db.PacketWriter.WriteLong(data, value, index)
}

func (db *DefaultBuffer) WriteULong(data []byte, value uint64, index uint64) {
	db.Next(8)
	db.PacketWriter.WriteULong(data, value, index)
}

func (db *DefaultBuffer) WriteShort(data []byte, value int16, index uint64) {
	db.Next(2)
	db.PacketWriter.WriteShort(data, value, index)
}

func (db *DefaultBuffer) WriteUShort(data []byte, value uint16, index uint64) {
	db.Next(2)
	db.PacketWriter.WriteUShort(data, value, index)
}

func (db *DefaultBuffer) WriteString(data []byte, value string, index uint64) {
	db.Next(uint64(len(value)))
	db.PacketWriter.WriteString(data, value, index)
}

func (db *DefaultBuffer) WriteBytes(data []byte, value []byte, index uint64) {
	db.Next(uint64(len(value)))
	db.PacketWriter.WriteBytes(data, value, index)
}

func (db *DefaultBuffer) ReadBool(data []byte, index uint64) bool {
	db.Next(1)
	return db.PacketReader.ReadBool(data, index)
}

func (db *DefaultBuffer) ReadByte(data []byte, index uint64) byte {
	db.Next(1)
	return db.PacketReader.ReadByte(data, index)
}

func (db *DefaultBuffer) ReadFloat(data []byte, index uint64) float32 {
	db.Next(4)
	return db.PacketReader.ReadFloat(data, index)
}

func (db *DefaultBuffer) ReadDouble(data []byte, index uint64) float64 {
	db.Next(8)
	return db.PacketReader.ReadDouble(data, index)
}

func (db *DefaultBuffer) ReadInt(data []byte, index uint64) int32 {
	db.Next(4)
	return db.PacketReader.ReadInt(data, index)
}

func (db *DefaultBuffer) ReadUint(data []byte, index uint64) uint32 {
	db.Next(4)
	return db.PacketReader.ReadUint(data, index)
}

func (db *DefaultBuffer) ReadLong(data []byte, index uint64) int64 {
	db.Next(8)
	return db.PacketReader.ReadLong(data, index)
}

func (db *DefaultBuffer) ReadULong(data []byte, index uint64) uint64 {
	db.Next(8)
	return db.PacketReader.ReadULong(data, index)
}

func (db *DefaultBuffer) ReadShort(data []byte, index uint64) int16 {
	db.Next(2)
	return db.PacketReader.ReadShort(data, index)
}

func (db *DefaultBuffer) ReadUShort(data []byte, index uint64) uint16 {
	db.Next(2)
	return db.PacketReader.ReadUShort(data, index)
}

func (db *DefaultBuffer) ReadString(data []byte, index uint64, length uint64) string {
	db.Next(length)
	return db.PacketReader.ReadString(data, index, length)
}

func (db *DefaultBuffer) ReadBytes(data []byte, index uint64, length uint64) []byte {
	db.Next(length)
	return db.PacketReader.ReadBytes(data, index, length)
}

func (db *DefaultBuffer) Clone() PacketBuffer {
	return &DefaultBuffer{
		PacketReader: db.PacketReader,
		PacketWriter: db.PacketWriter,
		initLen:      db.initLen,
		maxLen:       db.maxLen,
		buf:          make([]byte, db.maxLen),
	}
}

func (db *DefaultBuffer) SetInitLength(length uint64) {
	db.initLen = length
}

func (db *DefaultBuffer) SetMaxLength(size uint64) {
	db.maxLen = size
	db.buf = make([]byte, size)
}

func (db *DefaultBuffer) SetReader(reader PacketReader) {
	db.PacketReader = reader
}

func (db *DefaultBuffer) SetWriter(writer PacketWriter) {
	db.PacketWriter = writer
}

func (db *DefaultBuffer) Resize(size uint64) {
	if size > db.maxLen {
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
	if amount > db.index {
		db.index = 0
	} else {
		db.index -= amount
	}
	return db.index
}

func (db *DefaultBuffer) Reset() {
	db.index = 0
	db.len = db.initLen
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
