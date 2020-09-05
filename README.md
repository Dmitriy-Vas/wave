# Wave
Wave provides configurable MITM proxy to read and edit TCP packets
![](wave_gopher.png)

## Table of Content
+ [Features](https://github.com/Dmitriy-Vas/wave/blob/master/README.md#features)
+ [Getting Started](https://github.com/Dmitriy-Vas/wave/blob/master/README.md#getting-started)
+ [Contributing](https://github.com/Dmitriy-Vas/wave/blob/master/README.md#contributing)

## Features
+ Hook packets by their ID
+ Edit packets on-fly
+ Prevent unnecessary packets from being sent
+ Manually create and send custom packet

## Getting started
You can find full example on [cmd/proxy](https://github.com/Dmitriy-Vas/wave/blob/master/cmd/proxy/main.go)

Import the `wave` package from GitHub in your code:
```go
import "github.com/Dmitriy-Vas/wave"
```

Create remote and local addresses, they must implement `net.Addr` interface.
```go
localAddr, _ := net.ResolveTCPAddr("tcp", ":7999")
remoteAddr, _ := net.ResolveTCPAddr("tcp", "74.91.123.86:7000")
```

Create buffer, must implement `wave.PacketBuffer` interface.
You can use default buffer, but I'd recommend create your own wrapper.
```go
buf := &buffer.DefaultBuffer{
    PacketReader: &buffer.DefaultReader{Order: binary.LittleEndian},
    PacketWriter: &buffer.DefaultWriter{Order: binary.LittleEndian},
}
```

Create `wave.Config` and fill with your parameters
```go
config := wave.Config{
    // First bytes of packet, means packet payload size
    PacketLengthSize:   wave.Size32Bits,
    // Packet type usually goes right after payload size
    PacketTypeSize:     wave.Size8Bits,
    // Sometimes PacketLengthSize may include self bytes, in this example it will be additional 32bits (4 bytes)
    LengthIncludesSelf: false,

    // Remote address of server, wave will connect to this address
    RemoteAddress:      remoteAddr,
    // Local address, wave will listen for new connection on this address
    LocalAddress:       localAddr,
    // If you want to establish connection to the remove server right after receiving local connection
    ConnectImmediately: true,

    // Implementation of wave.PacketBuffer
    Buffer:             buf,

    // In current example packet is encrypted using RC4
    // I want to decrypt on start (to read) and encrypt on end (to send)
    OutgoingProcess: func(buffer buffer.PacketBuffer, start bool) {
        raw, _ := hex.DecodeString("c79332b197f92ba85ed281a023")
        cipher, _ := rc4.NewCipher(raw)
        cipher.XORKeyStream(buffer.Bytes()[5:], buffer.Bytes()[5:])
    },
    IncomingProcess: func(buffer buffer.PacketBuffer, start bool) {
        raw, _ := hex.DecodeString("6a39570cc9de4ec71d64821894")
        cipher, _ := rc4.NewCipher(raw)
        cipher.XORKeyStream(buffer.Bytes()[5:], buffer.Bytes()[5:])
    },
}
```

Create new `wave.Proxy` instance
```go
proxy := wave.New(config)
```

If you want to edit packets, you must create packet structure and register to `wave.Proxy`.
Check [lib/packets](https://github.com/Dmitriy-Vas/wave/tree/master/lib/packets) for examples.
```go
// 1 is packet ID
// true means this packet is outgoing (from client to server)
// last parameter is a pointer to the packet, must implement wave.Packet interface
proxy.AddPacket(1, true, new(outgoing.NewAccountPacket))
```

To actually edit packet, register your hook and do magic
```go
proxy.HookPacket(int64(lib.IncReceiveHour), false, func(conn *wave.Conn, packet wave.Packet) {
    receiveHourPacket := packet.(*incoming.ReceiveHourPacket)
    log.Printf("Changing hour from %d to 12", receiveHourPacket.Hour)
    receiveHourPacket.Hour = 12
})
```

### Contributing

If you want to contribute, fork this project, commit changes and create pull request.
Please describe your changes and what they are doing in pull request.