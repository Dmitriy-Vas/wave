package wave

import (
	"net"
	"reflect"
	"sync"
	"sync/atomic"
)

type Proxy struct {
	listener                             net.Listener // Local listener, listening on LocalAddress for new connection
	index                                *uint32      // Amount of served local connections
	Connections                          *sync.Map    // map[uint32]*Conn. Concurrency safe. Map with all current connections
	ClientPacketHooks, ServerPacketHooks *sync.Map    // map[id][]PacketHook Concurrency safe. Map with all registered packet hooks
	ClientPacketMap, ServerPacketMap     *sync.Map    // map[id]Packet. Concurrency safe. Map with all available packets
	Config                               Config
}

type PacketHook func(conn *Conn, packet Packet)

// Start listening on LocalAddress and accept new connections
// Every connection have their unique ID, starting from 0
// Connections are stored by unique ID inside Proxy.Connections
// Note: blocks execution
func (p *Proxy) Start() (err error) {
	if p.listener, err = net.Listen("tcp", p.Config.LocalAddress.String()); err != nil {
		return err
	}

	for {
		local_conn, err := p.listener.Accept()
		if err != nil {
			return err
		}
		id := atomic.AddUint32(p.index, 1)

		conn := &Conn{
			ID:        id,
			LocalConn: local_conn,
			Done:      make(chan error),
			proxy:     p,
		}
		p.Connections.Store(id, conn)

		go conn.Start()
	}
}

func (p *Proxy) Index() uint32 {
	return atomic.LoadUint32(p.index)
}

// Closes every connection and stops the listener
func (p *Proxy) Close() {
	p.Connections.Range(func(key, value interface{}) bool {
		if connection, ok := value.(*Conn); ok {
			connection.Close(nil)
		}
		return true
	})
	p.listener.Close()
}

// Adds function, which will be fired on every packet with specified ID
// Outgoing=true, if your packet is going from client to server
// Outgoing=false, if your packet is going from server to client
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

// Adds Packet to the list of available packets, 1 packet per 1 ID
// Outgoing=true, if your packet is going from client to server
// Outgoing=false, if your packet is going from server to client
func (p *Proxy) AddPacket(ID int64, outgoing bool, packet Packet) {
	packetType := reflect.TypeOf(packet)
	if outgoing {
		p.ClientPacketMap.Store(ID, packetType)
	} else {
		p.ServerPacketMap.Store(ID, packetType)
	}
}

// Creates new proxy instance with specified config
func New(config Config) *Proxy {
	return &Proxy{
		index:             new(uint32),
		Connections:       new(sync.Map),
		ClientPacketHooks: new(sync.Map),
		ServerPacketHooks: new(sync.Map),
		ClientPacketMap:   new(sync.Map),
		ServerPacketMap:   new(sync.Map),
		Config:            config,
	}
}
