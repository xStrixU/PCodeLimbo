package protocol

import (
	"encoding/binary"
	"fmt"
	"io"
)

type Writer struct {
	w interface {
		io.Writer
		io.ByteWriter
	}
}

func NewWriter(w interface {
	io.Writer
	io.ByteWriter
}) * Writer {
	return &Writer{w}
}

func (writer *Writer) Write(v interface{}) {
	if err := binary.Write(writer.w, binary.BigEndian, v); err != nil {
		fmt.Println("Protocol write error: ", err)
	}
}

func (writer *Writer) WriteShort(v int16) {
	writer.Write(v)
}

func (writer *Writer) WriteVarInt(v VarInt) {
	var buf [8]byte
	n := binary.PutUvarint(buf[:], uint64(uint32(v)))

	if _, err := writer.w.Write(buf[:n]); err != nil {
		fmt.Println("Protocol write VarInt error: ", err)
	}
}

func (writer *Writer) WriteString(s string) {
	writer.WriteVarInt(VarInt(len(s)))

	if _, err := writer.w.Write([]byte(s)); err != nil {
		fmt.Println("Protocol write String error: ", err)
	}
}