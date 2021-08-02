package protocol

import (
	"encoding/binary"
	"fmt"
	"io"
)

type Reader struct {
	r interface {
		io.Reader
		io.ByteReader
	}
}

func NewReader(r interface {
	io.Reader
	io.ByteReader
}) *Reader {
	return &Reader{r}
}

func (reader *Reader) Read(v interface{}) {
	if err := binary.Read(reader.r, binary.BigEndian, v); err != nil {
		fmt.Println("Protocol read error: ", err)
	}
}

func (reader *Reader) ReadShort() (v int16) {
	reader.Read(&v)
	return
}

func (reader *Reader) ReadVarInt() VarInt {
	v, err := binary.ReadUvarint(reader.r)

	if err != nil {
		fmt.Println("Protocol read VarInt error: ", err)
		return 0
	}

	return VarInt(int32(uint32(v)))
}

func (reader *Reader) ReadString() string {
	length := reader.ReadVarInt()
	buf := make([]byte, length)

	if _, err := io.ReadFull(reader.r, buf); err != nil {
		fmt.Println("Protocol read String error: ", err)
		return ""
	}

	return string(buf)
}