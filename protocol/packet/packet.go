package packet

import (
	"bytes"
	"github.com/xStrixU/PCodeLimbo/protocol"
)

const (
	IDHandshakeClientBound = protocol.VarInt(0x00)
	IDHandshakeServerBound = protocol.VarInt(0x00)
)

const (
	NextStateStatus = protocol.VarInt(1)
	NextStateLogin = protocol.VarInt(2)
)

type Packet interface {
	ID() protocol.VarInt
	encode(*protocol.Writer)
	Decode(*protocol.Reader) // this is public because this func is using in Player#HandlePacket()
}

func EncodePacket(packet Packet) []byte {
	buf := &bytes.Buffer{}
	writer := protocol.NewWriter(buf)

	packet.encode(writer)
	packetLen := buf.Len()
	buf.Reset()

	writer.WriteVarInt(protocol.VarInt(packetLen + protocol.VarIntSize(int(packet.ID())))) // packet length
	writer.WriteVarInt(packet.ID())
	packet.encode(writer)

	return buf.Bytes()
}