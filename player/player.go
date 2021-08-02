package player

import (
	"bytes"
	"fmt"
	"github.com/xStrixU/PCodeLimbo/protocol"
	"github.com/xStrixU/PCodeLimbo/protocol/packet"
	"net"
)

type Player struct {
	Name string
	NextState protocol.VarInt
	conn net.Conn
}

func NewPlayer(conn net.Conn) *Player {
	return &Player{
		Name: "",
		NextState: packet.NextStateStatus,
		conn: conn,
	}
}

func (p *Player) Start() {
	buf := make([]byte, 1024)
	length, err := p.conn.Read(buf)

	if err != nil {
		fmt.Println("Read packet error: ", err)
	}

	p.HandlePacket(buf[:length])
}

func (p *Player) HandlePacket(data []byte) {
	buf := bytes.NewBuffer(data)
	reader := protocol.NewReader(buf)

	reader.ReadVarInt() // packetLength, TODO: Handle it
	packetID := reader.ReadVarInt()

	if p.NextState == packet.NextStateStatus && packetID == packet.IDHandshakeServerBound {
		json := "{\"description\":{\"text\":\"opis serwera\"},\"players\":{\"max\":100,\"online\":0},\"version\":{\"name\":\"nazwa silnika\",\"protocol\":47}}" // TODO: change it

		response := &packet.HandshakeClientBound{Json: json}
		packet.EncodePacket(response)
		p.SendPacket(response)

		handshakeServerBound := &packet.HandshakeServerBound{}
		handshakeServerBound.Decode(reader)

		p.NextState = handshakeServerBound.NextState
	}
}

func (p *Player) SendPacket(pk packet.Packet) {
	data := packet.EncodePacket(pk)

	if _, err := p.conn.Write(data); err != nil {
		fmt.Println("Send packet to player error: ", err)
	}
}