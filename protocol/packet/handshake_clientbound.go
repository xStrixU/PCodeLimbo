package packet

import "github.com/xStrixU/PCodeLimbo/protocol"

type HandshakeClientBound struct {
	Json string
}

func (packet *HandshakeClientBound) ID() protocol.VarInt {
	return IDHandshakeServerBound
}

func (packet *HandshakeClientBound) encode(writer *protocol.Writer) {
	writer.WriteString(packet.Json)
}

func (packet *HandshakeClientBound) Decode(*protocol.Reader) {
	// don't need to decode this packet
}