package wave

import (
	. "github.com/Dmitriy-Vas/wave/buffer"
	. "github.com/Dmitriy-Vas/wave/packets"
	"log"
	"net"
	"sync"
	"sync/atomic"
)

type Proxy struct {
	RemoteAddr                           net.Addr
	Listener                             *net.Listener
	_buffer                              PacketBuffer
	_reader                              PacketReader
	_writer                              PacketWriter
	_buffer_args                         []interface{}
	Index                                *uint32   // Amount of served local connections
	Connections                          *sync.Map // map[uint32]*Conn
	ClientPacketHooks, ServerPacketHooks *sync.Map // MapHooks
	ClientPacketMap, ServerPacketMap     *sync.Map // MapPackets
}

type PacketHook func(conn *Conn, packet Packet)
type MapHooks map[int64][]PacketHook
type MapPackets map[int64]Packet

// Note: blocking execution
func (p *Proxy) Start(LocalAddr net.Addr) error {
	listener, err := net.Listen("tcp", LocalAddr.String())
	if err != nil {
		return err
	}

	for {
		local_conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		id := atomic.AddUint32(p.Index, 1)

		conn := &Conn{
			ID:          id,
			LocalConn:   local_conn,
			Done:        make(chan error),
			_proxy:      p,
			clientHooks: new(sync.Map),
			serverHooks: new(sync.Map),
		}
		p.Connections.Store(id, conn)

		go conn.Start()
	}
}

func (p *Proxy) Close() {
	p.Connections.Range(func(key, value interface{}) bool {
		if connection, ok := value.(*Conn); ok {
			connection.Close(nil)
		}
		return true
	})
}

func (p *Proxy) HookPacket(ID int64, outgoing bool, hook PacketHook) {
	var m *sync.Map
	if outgoing {
		m = p.ClientPacketHooks
	} else {
		m = p.ServerPacketHooks
	}
	hooks := make([]PacketHook, 0)
	hooks = append(hooks, hook)
	if _hooks, ok := m.LoadOrStore(ID, hooks); ok {
		hooks = _hooks.([]PacketHook)
		hooks = append(hooks, hook)
		m.Store(ID, hooks)
	}
}

func (p *Proxy) AddPacket(ID int64, outgoing bool, packet Packet) {
	if outgoing {
		p.ClientPacketMap.Store(ID, packet)
	} else {
		p.ServerPacketMap.Store(ID, packet)
	}
}

func (p *Proxy) SetBuffer(buffer PacketBuffer, args ...interface{}) {
	p._buffer = buffer
	p._buffer_args = args
}

func (p *Proxy) SetReader(reader PacketReader) {
	p._reader = reader
}

func (p *Proxy) SetWriter(writer PacketWriter) {
	p._writer = writer
}

func New(RemoteAddr net.Addr) *Proxy {
	return &Proxy{
		RemoteAddr:        RemoteAddr,
		_buffer:           (*DefaultBuffer)(nil),
		_reader:           (*DefaultReader)(nil),
		_writer:           (*DefaultWriter)(nil),
		_buffer_args:      []interface{}{LengthSize},
		Index:             new(uint32),
		Connections:       new(sync.Map),
		ClientPacketHooks: new(sync.Map),
		ServerPacketHooks: new(sync.Map),
		ClientPacketMap:   new(sync.Map),
		ServerPacketMap:   new(sync.Map),
	}
}
