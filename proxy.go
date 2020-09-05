package wave

import (
	"net"
	"reflect"
	"sync"
	"sync/atomic"
)

type Proxy struct {
	listener                                 net.Listener // Local listener, listening on LocalAddress for new connection
	index                                    *uint32      // Amount of served local connections
	Connections                              *sync.Map    // map[uint32]*Conn. Concurrency safe. Map with all current connections
	OutgoingPacketHooks, IncomingPacketHooks *sync.Map    // map[id][]PacketHook Concurrency safe. Map with all registered packet hooks
	OutgoingPacketMap, IncomingPacketMap     *sync.Map    // map[id]Packet. Concurrency safe. Map with all available packets
	Config                                   Config       // Proxy configuration
}

// PacketHook is a shortcut for packet hook function.
// Conn is a connection, which are received packet.
// Packet is an interface Packet, can be asserted to the actual packet type.
type PacketHook func(conn *Conn, packet Packet)

// Start starts listening on LocalAddress and accept new connections.
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
		// Store connection inside map
		p.Connections.Store(id, conn)

		go conn.start()
	}
}

// Index returns current amount of served connections.
func (p *Proxy) Index() uint32 {
	return atomic.LoadUint32(p.index)
}

// Close closes every connection and stops the listener.
func (p *Proxy) Close() {
	p.Connections.Range(func(key, value interface{}) bool {
		if connection, ok := value.(*Conn); ok {
			connection.Close(nil)
		}
		return true
	})
	p.listener.Close()
}

// HookPacket adds function, which will be fired for every packet with specified ID.
// Outgoing=true, if your packet is going from client to server.
// Outgoing=false, if your packet is going from server to client.
func (p *Proxy) HookPacket(ID int64, outgoing bool, hook PacketHook) {
	var m *sync.Map
	if outgoing {
		m = p.OutgoingPacketHooks
	} else {
		m = p.IncomingPacketHooks
	}
	hooks := make([]PacketHook, 0)
	hooks = append(hooks, hook)
	if _hooks, ok := m.LoadOrStore(ID, hooks); ok {
		hooks = _hooks.([]PacketHook)
		hooks = append(hooks, hook)
		m.Store(ID, hooks)
	}
}

// AddPacket adds Packet to the list of available to hook packets.
// Outgoing=true, if your packet is going from client to server.
// Outgoing=false, if your packet is going from server to client.
func (p *Proxy) AddPacket(ID int64, outgoing bool, packet Packet) {
	packetType := reflect.TypeOf(packet)
	if outgoing {
		p.OutgoingPacketMap.Store(ID, packetType)
	} else {
		p.IncomingPacketMap.Store(ID, packetType)
	}
}

// New creates new proxy instance with specified config.
func New(config Config) *Proxy {
	return &Proxy{
		index:               new(uint32),
		Connections:         new(sync.Map),
		OutgoingPacketHooks: new(sync.Map),
		IncomingPacketHooks: new(sync.Map),
		OutgoingPacketMap:   new(sync.Map),
		IncomingPacketMap:   new(sync.Map),
		Config:              config,
	}
}
