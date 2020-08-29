package buffer

import (
	"sync/atomic"
)

type PacketBuffer interface {
	PacketReader
	PacketWriter
	SetReader(reader PacketReader)
	SetWriter(writer PacketWriter)
	Init(args []interface{})

	AppendIfNeeded(bytes uint32)
	Resize(size uint32, index uint32)
	Reindex()
	Next(amount uint32) uint32
	Reset()
	Len() uint32
	Index() uint32
	Bytes() []byte
}

type Vector2 struct {
	First  int32
	Second int32
}

// TODO rewrite to zero-copy buffer implementation
type DefaultBuffer struct {
	PacketReader
	PacketWriter

	init_len uint32
	buf      []byte
	index    *uint32
}

func (db *DefaultBuffer) AppendIfNeeded(bytes uint32) {
	remaining := db.Len() - db.Index()
	if remaining < bytes {
		db.buf = append(db.buf, make([]byte, 0, bytes-remaining)...)
	}
}

const (
	DefaultLength = 8
)

// TODO add better solution to pass arguments
func (db *DefaultBuffer) Init(args []interface{}) {
	var init_len uint32 = DefaultLength
	db.init_len = init_len
	db.buf = make([]byte, init_len)
	db.index = new(uint32)
}

func (db *DefaultBuffer) SetReader(reader PacketReader) {
	db.PacketReader = reader
}

func (db *DefaultBuffer) SetWriter(writer PacketWriter) {
	db.PacketWriter = writer
}

func (db *DefaultBuffer) Resize(size uint32, index uint32) {
	if size > 0xfffff {
		panic("New buffer size is too large")
	}
	old := db.buf
	buf := make([]byte, size)
	for i, b := range old {
		buf[int(index)+i] = b
	}
	db.buf = buf
}

func (db *DefaultBuffer) Reindex() {
	atomic.StoreUint32(db.index, 0)
}

func (db *DefaultBuffer) Next(amount uint32) uint32 {
	return atomic.AddUint32(db.index, amount)
}

func (db *DefaultBuffer) Reset() {
	db.buf = make([]byte, db.init_len)
	atomic.StoreUint32(db.index, 0)
}

func (db *DefaultBuffer) Len() uint32 {
	return uint32(len(db.buf))
}

func (db *DefaultBuffer) Index() uint32 {
	return atomic.LoadUint32(db.index)
}

func (db *DefaultBuffer) Bytes() []byte {
	return db.buf
}
