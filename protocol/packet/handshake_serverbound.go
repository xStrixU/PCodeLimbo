package packet

import "github.com/xStrixU/PCodeLimbo/protocol"

type HandshakeServerBound struct {
	ProtocolVersion protocol.VarInt
	Address string
	Port int16
	NextState protocol.VarInt
}

func (packet *HandshakeServerBound) ID() protocol.VarInt {
	return IDHandshakeClientBound
}

func (packet *HandshakeServerBound) encode(*protocol.Writer) {
	// don't need to encode this packet
}

func (packet *HandshakeServerBound) Decode(reader *protocol.Reader) {
	packet.ProtocolVersion = reader.ReadVarInt()
	packet.Address = reader.ReadString()
	packet.Port = reader.ReadShort()
	packet.NextState = reader.ReadVarInt()
}